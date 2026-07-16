package snbgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	config     *Config
	httpClient *http.Client
	token      *TokenManager
	idGen      *OrderIDGenerator
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Millisecond,
		},
		token: &TokenManager{},
		idGen: &OrderIDGenerator{},
	}
}

func (c *Client) IDGenerator() *OrderIDGenerator {
	return c.idGen
}

func (c *Client) Login() (*APIResponse, error) {
	req := buildLoginRequest(c.config.Account, c.config.Key)
	resp, err := c.execute(req)
	if err != nil {
		return nil, err
	}
	if resp.Succeed() {
		var data TokenData
		if err := resp.UnmarshalData(&data); err == nil {
			c.token.Set(data.AccessToken, data.ExpiryTime)
		}
	}
	return resp, nil
}

func (c *Client) GetTokenStatus() (*APIResponse, error) {
	token, ok := c.token.Get()
	if !ok {
		return nil, ErrLoginNeeded
	}
	req := buildGetTokenStatusRequest(c.config.Account, token)
	return c.execute(req)
}

func (c *Client) PlaceOrder(orderID string, securityType SecurityType, symbol, exchange string, side OrderSide, currency Currency, quantity int, price float64, orderType OrderType, tif TimeInForce, forceOnlyRTH bool, stopPrice float64, parent string, orderIDType OrderIdType, tradingHours TradingHours) (*APIResponse, error) {
	params := map[string]string{
		"security_type": string(securityType),
		"symbol":        symbol,
		"exchange":      exchange,
		"side":          string(side),
		"currency":      string(currency),
		"quantity":      itoa(quantity),
		"price":         fmt.Sprintf("%g", price),
		"order_type":    string(orderType),
		"tif":           string(tif),
		"rth":           boolStr(forceOnlyRTH),
		"order_id_type": string(orderIDType),
	}
	if stopPrice > 0 {
		params["stop_price"] = fmt.Sprintf("%g", stopPrice)
	}
	if parent != "" {
		params["parent"] = parent
	}
	if tradingHours != "" {
		params["trading_hours"] = string(tradingHours)
	}
	req := buildPlaceOrderRequest(c.config.Account, orderID, params)
	return c.execute(req)
}

func (c *Client) CancelOrder(orderID, originOrderID string, orderIDType OrderIdType) (*APIResponse, error) {
	req := buildCancelOrderRequest(c.config.Account, orderID, originOrderID, orderIDType)
	return c.execute(req)
}

func (c *Client) GetOrderByID(orderID string) (*APIResponse, error) {
	req := buildGetOrderByIDRequest(c.config.Account, orderID)
	return c.execute(req)
}

func (c *Client) GetOrderList(page, size int, status, securityType string) (*APIResponse, error) {
	req := buildGetOrderListRequest(c.config.Account, page, size, status, securityType)
	return c.execute(req)
}

func (c *Client) GetBalance() (*APIResponse, error) {
	req := buildGetBalanceRequest(c.config.Account)
	return c.execute(req)
}

func (c *Client) GetPositionList(securityType string) (*APIResponse, error) {
	req := buildGetPositionListRequest(c.config.Account, securityType)
	return c.execute(req)
}

func (c *Client) GetTransactionList(page, size int, side OrderSide, orderTimeMin, orderTimeMax int64) (*APIResponse, error) {
	req := buildGetTransactionListRequest(c.config.Account, page, size, side, orderTimeMin, orderTimeMax)
	return c.execute(req)
}

func (c *Client) GetSecurityDetail(symbol string) (*APIResponse, error) {
	req := buildGetSecurityDetailRequest(c.config.Account, symbol)
	return c.execute(req)
}

func (c *Client) execute(req httpRequest) (*APIResponse, error) {
	if req.auth {
		token, ok := c.token.Get()
		if !ok {
			loginResp, err := c.Login()
			if err != nil {
				return nil, fmt.Errorf("snbgo: auto-login failed: %w", err)
			}
			if !loginResp.Succeed() {
				return loginResp, nil
			}
			token, _ = c.token.Get()
		}
		if req.params == nil {
			req.params = make(map[string]string)
		}
		// Token is sent via Cookie header, not in params
		_ = token
	}

	fullURL := c.config.BaseURL() + "/" + req.path

	var body io.Reader
	headers := map[string]string{
		"User-Agent":   "snbgo/1.0.0",
		"Accept":       "application/vnd.snowx+json; version=1.0",
		"Cache-Control": "no-cache",
		"Connection":   "Keep-Alive",
	}

	if req.auth {
		token, _ := c.token.Get()
		headers["Cookie"] = "access_token=" + token
	}

	switch req.method {
	case "GET", "DELETE":
		if len(req.params) > 0 {
			fullURL += "?" + encodeParams(req.params)
		}
	case "POST":
		body = strings.NewReader(encodeParams(req.params))
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	httpReq, err := http.NewRequest(req.method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("snbgo: failed to create request: %w", err)
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("snbgo: HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("snbgo: failed to read response: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("snbgo: failed to parse response: %w (body: %s)", err, string(respBody))
	}

	return &apiResp, nil
}

func encodeParams(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	return values.Encode()
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
