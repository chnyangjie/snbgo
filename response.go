package snbgo

import "encoding/json"

type APIResponse struct {
	ResultCode string          `json:"result_code"`
	Msg        json.RawMessage `json:"msg"`
	ResultData json.RawMessage `json:"result_data"`
}

func (r *APIResponse) Succeed() bool {
	return r.ResultCode == "60000"
}

func (r *APIResponse) Message() string {
	if r.Msg == nil {
		return ""
	}
	var s string
	if err := json.Unmarshal(r.Msg, &s); err == nil {
		return s
	}
	return string(r.Msg)
}

func (r *APIResponse) DataString() string {
	if r.ResultData == nil {
		return "null"
	}
	return string(r.ResultData)
}

func (r *APIResponse) UnmarshalData(v interface{}) error {
	return json.Unmarshal(r.ResultData, v)
}

type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiryTime  int64  `json:"expiry_time"`
}

type OrderData struct {
	AccountID        string   `json:"account_id"`
	AveragePrice     float64  `json:"average_price"`
	Children         *string  `json:"children"`
	ContractID       *string  `json:"contract_id"`
	Currency         string   `json:"currency"`
	Exchange         string   `json:"exchange"`
	FilledQuantity   int      `json:"filled_quantity"`
	GroupID          *string  `json:"group_id"`
	ID               string   `json:"id"`
	IPAddress        *string  `json:"ip_address"`
	LastFilledTime   int64    `json:"last_filled_time"`
	Memo             string   `json:"memo"`
	OrderTime        int64    `json:"order_time"`
	OrderType        string   `json:"order_type"`
	Parent           *string  `json:"parent"`
	Price            float64  `json:"price"`
	Quantity         int      `json:"quantity"`
	RTH              bool     `json:"rth"`
	SecondaryOrderID string   `json:"secondary_order_id"`
	SecurityID       *string  `json:"security_id"`
	SecurityIDType   *string  `json:"security_id_type"`
	SecurityType     string   `json:"security_type"`
	Side             string   `json:"side"`
	SNBOrderID       string   `json:"snb_order_id"`
	Status           string   `json:"status"`
	StopPrice        *string  `json:"stop_price"`
	Symbol           string   `json:"symbol"`
	TIF              string   `json:"tif"`
	TradingHours     string   `json:"trading_hours"`
}

type OrderListData struct {
	Count int         `json:"count"`
	Items []OrderData `json:"items"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

type PlaceOrderResult struct {
	ID         string `json:"id"`
	Memo       string `json:"memo"`
	SNBOrderID string `json:"snb_order_id"`
	Status     string `json:"status"`
}

type CancelOrderResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type PositionData struct {
	AccountID    string  `json:"account_id"`
	AveragePrice float64 `json:"average_price"`
	ContractID   string  `json:"contract_id"`
	Exchange     string  `json:"exchange"`
	MarketPrice  float64 `json:"market_price"`
	Position     float64 `json:"position"`
	RealizedPnL  float64 `json:"realized_pnl"`
	SecurityType string  `json:"security_type"`
	Symbol       string  `json:"symbol"`
}

type BalanceDetailItem struct {
	Cash     float64 `json:"cash"`
	Currency string  `json:"currency"`
}

type BalanceData struct {
	BalanceDetailItems             []BalanceDetailItem `json:"balance_detail_items"`
	Cash                           float64             `json:"cash"`
	Currency                       string              `json:"currency"`
	CurrentAvailableFunds          float64             `json:"current_available_funds"`
	CurrentExcessLiquidity         float64             `json:"current_excess_liquidity"`
	CurrentInitialMargin           float64             `json:"current_initial_margin"`
	CurrentMaintenanceMargin       float64             `json:"current_maintenance_margin"`
	EquityWithLoanValue            float64             `json:"equity_with_loan_value"`
	Leverage                       float64             `json:"leverage"`
	NetLiquidationValue            float64             `json:"net_liquidation_value"`
	PreviousDayEquityWithLoanValue float64             `json:"previous_day_equity_with_loan_value"`
	SecuritiesGrossPositionValue   float64             `json:"securities_gross_position_value"`
	SMA                            float64             `json:"sma"`
}

type TransactionData struct {
	AccountID      string  `json:"account_id"`
	Currency       string  `json:"currency"`
	Exchange       string  `json:"exchange"`
	ID             string  `json:"id"`
	OrderPrice     float64 `json:"order_price"`
	OrderQuantity  float64 `json:"order_quantity"`
	OrderTime      int64   `json:"order_time"`
	OrderType      string  `json:"order_type"`
	Price          float64 `json:"price"`
	Quantity       float64 `json:"quantity"`
	RTH            bool    `json:"rth"`
	SNBOrderID     string  `json:"snb_order_id"`
	SecurityType   string  `json:"security_type"`
	Side           string  `json:"side"`
	Status         string  `json:"status"`
	Symbol         string  `json:"symbol"`
	TIF            string  `json:"tif"`
	TradeTime      int64   `json:"trade_time"`
}

type TransactionListData struct {
	Count int               `json:"count"`
	Items []TransactionData `json:"items"`
	Page  int               `json:"page"`
	Size  int               `json:"size"`
}

type SecurityDetailItem struct {
	Exchange     string `json:"exchange"`
	LotSize      string `json:"lot_size"`
	SecurityName string `json:"security_name"`
	Symbol       string `json:"symbol"`
	TickSize     string `json:"tick_size"`
}
