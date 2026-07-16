package snbgo

import "fmt"

type httpRequest struct {
	method string
	path   string
	params map[string]string
	auth   bool
}

func buildLoginRequest(accountID, secretKey string) httpRequest {
	return httpRequest{
		method: "POST",
		path:   "auth/" + accountID + "/access-token",
		params: map[string]string{"secret_key": secretKey},
		auth:   false,
	}
}

func buildGetTokenStatusRequest(accountID, token string) httpRequest {
	return httpRequest{
		method: "GET",
		path:   "/auth/" + accountID + "/access-token/" + token,
		params: nil,
		auth:   true,
	}
}

func buildGetOrderListRequest(accountID string, page, size int, status, securityType string) httpRequest {
	params := map[string]string{
		"account_id": accountID,
		"page":       itoa(page),
		"size":       itoa(size),
	}
	if status != "" {
		params["status"] = status
	}
	if securityType != "" {
		params["security_type"] = securityType
	}
	return httpRequest{
		method: "GET",
		path:   "order",
		params: params,
		auth:   true,
	}
}

func buildGetOrderByIDRequest(accountID, orderID string) httpRequest {
	return httpRequest{
		method: "GET",
		path:   "order/" + orderID,
		params: map[string]string{"account_id": accountID},
		auth:   true,
	}
}

func buildPlaceOrderRequest(accountID, orderID string, params map[string]string) httpRequest {
	params["account_id"] = accountID
	return httpRequest{
		method: "POST",
		path:   "order/" + orderID,
		params: params,
		auth:   true,
	}
}

func buildCancelOrderRequest(accountID, orderID, originOrderID string, orderIDType OrderIdType) httpRequest {
	return httpRequest{
		method: "DELETE",
		path:   "order/" + originOrderID,
		params: map[string]string{
			"account_id":    accountID,
			"new_id":        orderID,
			"order_id_type": string(orderIDType),
		},
		auth: true,
	}
}

func buildGetBalanceRequest(accountID string) httpRequest {
	return httpRequest{
		method: "GET",
		path:   "funds",
		params: map[string]string{"account_id": accountID},
		auth:   true,
	}
}

func buildGetPositionListRequest(accountID, securityType string) httpRequest {
	params := map[string]string{"account_id": accountID}
	if securityType != "" {
		params["security_type"] = securityType
	}
	return httpRequest{
		method: "GET",
		path:   "position",
		params: params,
		auth:   true,
	}
}

func buildGetTransactionListRequest(accountID string, page, size int, side OrderSide, orderTimeMin, orderTimeMax int64) httpRequest {
	params := map[string]string{
		"account_id": accountID,
		"page":       itoa(page),
		"size":       itoa(size),
	}
	if side != "" {
		params["side"] = string(side)
	}
	if orderTimeMin > 0 {
		params["order_time_min"] = fmt.Sprintf("%d", orderTimeMin)
	}
	if orderTimeMax > 0 {
		params["order_time_max"] = fmt.Sprintf("%d", orderTimeMax)
	}
	return httpRequest{
		method: "GET",
		path:   "trade",
		params: params,
		auth:   true,
	}
}

func buildGetSecurityDetailRequest(accountID, symbol string) httpRequest {
	return httpRequest{
		method: "GET",
		path:   "security/details",
		params: map[string]string{
			"account_id": accountID,
			"symbol":     symbol,
		},
		auth: true,
	}
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	if i < 0 {
		return "-1"
	}
	buf := [20]byte{}
	pos := len(buf)
	for i > 0 {
		pos--
		buf[pos] = byte('0' + i%10)
		i /= 10
	}
	return string(buf[pos:])
}
