package snbgo

import "fmt"

type Currency string

const (
	CurrencyUSD  Currency = "USD"
	CurrencyHKD  Currency = "HKD"
	CurrencyCNY  Currency = "CNY"
	CurrencyCNH  Currency = "CNH"
	CurrencyEUR  Currency = "EUR"
	CurrencyGBP  Currency = "GBP"
	CurrencyJPY  Currency = "JPY"
	CurrencyAUD  Currency = "AUD"
	CurrencyCAD  Currency = "CAD"
	CurrencyCHF  Currency = "CHF"
	CurrencySGD  Currency = "SGD"
	CurrencyNZD  Currency = "NZD"
	CurrencySEK  Currency = "SEK"
	CurrencyNOK  Currency = "NOK"
	CurrencyDKK  Currency = "DKK"
	CurrencyPLN  Currency = "PLN"
	CurrencyCZK  Currency = "CZK"
	CurrencyHUF  Currency = "HUF"
	CurrencyTRY  Currency = "TRY"
	CurrencyZAR  Currency = "ZAR"
	CurrencyILS  Currency = "ILS"
	CurrencyMXN  Currency = "MXN"
	CurrencyRUB  Currency = "RUB"
	CurrencyKRW  Currency = "KRW"
	CurrencyALL  Currency = "ALL"
	CurrencyUSX  Currency = "USX"
)

var validCurrencies = map[Currency]bool{
	CurrencyUSD: true, CurrencyHKD: true, CurrencyCNY: true, CurrencyCNH: true,
	CurrencyEUR: true, CurrencyGBP: true, CurrencyJPY: true, CurrencyAUD: true,
	CurrencyCAD: true, CurrencyCHF: true, CurrencySGD: true, CurrencyNZD: true,
	CurrencySEK: true, CurrencyNOK: true, CurrencyDKK: true, CurrencyPLN: true,
	CurrencyCZK: true, CurrencyHUF: true, CurrencyTRY: true, CurrencyZAR: true,
	CurrencyILS: true, CurrencyMXN: true, CurrencyRUB: true, CurrencyKRW: true,
	CurrencyALL: true, CurrencyUSX: true,
}

func ParseCurrency(s string) (Currency, error) {
	c := Currency(s)
	if !validCurrencies[c] {
		return "", fmt.Errorf("invalid currency: %s", s)
	}
	return c, nil
}

type SecurityType string

const (
	SecurityTypeSTK   SecurityType = "STK"
	SecurityTypeOPT   SecurityType = "OPT"
	SecurityTypeFUT   SecurityType = "FUT"
	SecurityTypeWAR   SecurityType = "WAR"
	SecurityTypeIOPT  SecurityType = "IOPT"
	SecurityTypeFOP   SecurityType = "FOP"
	SecurityTypeCMDTY SecurityType = "CMDTY"
	SecurityTypeCFD   SecurityType = "CFD"
	SecurityTypeFUND  SecurityType = "FUND"
	SecurityTypeBOND  SecurityType = "BOND"
	SecurityTypeCASH  SecurityType = "CASH"
	SecurityTypeMLEG  SecurityType = "MLEG"
	SecurityTypeALL   SecurityType = "ALL"
)

var validSecurityTypes = map[SecurityType]bool{
	SecurityTypeSTK: true, SecurityTypeOPT: true, SecurityTypeFUT: true,
	SecurityTypeWAR: true, SecurityTypeIOPT: true, SecurityTypeFOP: true,
	SecurityTypeCMDTY: true, SecurityTypeCFD: true, SecurityTypeFUND: true,
	SecurityTypeBOND: true, SecurityTypeCASH: true, SecurityTypeMLEG: true,
	SecurityTypeALL: true,
}

func ParseSecurityType(s string) (SecurityType, error) {
	st := SecurityType(s)
	if !validSecurityTypes[st] {
		return "", fmt.Errorf("invalid security_type: %s", s)
	}
	return st, nil
}

type OrderType string

const (
	OrderTypeLIMIT             OrderType = "LIMIT"
	OrderTypeMARKET            OrderType = "MARKET"
	OrderTypeAT                OrderType = "AT"
	OrderTypeATL               OrderType = "ATL"
	OrderTypeSSL               OrderType = "SSL"
	OrderTypeSEL               OrderType = "SEL"
	OrderTypeSTOP              OrderType = "STOP"
	OrderTypeSTOP_LIMIT        OrderType = "STOP_LIMIT"
	OrderTypeTRAIL             OrderType = "TRAIL"
	OrderTypeTRAIL_LIMIT       OrderType = "TRAIL_LIMIT"
	OrderTypeLIMIT_ON_OPENING  OrderType = "LIMIT_ON_OPENING"
	OrderTypeMARKET_ON_OPENING OrderType = "MARKET_ON_OPENING"
	OrderTypeLIMIT_ON_CLOSE    OrderType = "LIMIT_ON_CLOSE"
	OrderTypeMARKET_ON_CLOSE   OrderType = "MARKET_ON_CLOSE"
	OrderTypeLIMIT_IF_TOUCHED  OrderType = "LIMIT_IF_TOUCHED"
	OrderTypeMARKET_IF_TOUCHED OrderType = "MARKET_IF_TOUCHED"
)

var validOrderTypes = map[OrderType]bool{
	OrderTypeLIMIT: true, OrderTypeMARKET: true, OrderTypeAT: true,
	OrderTypeATL: true, OrderTypeSSL: true, OrderTypeSEL: true,
	OrderTypeSTOP: true, OrderTypeSTOP_LIMIT: true, OrderTypeTRAIL: true,
	OrderTypeTRAIL_LIMIT: true, OrderTypeLIMIT_ON_OPENING: true,
	OrderTypeMARKET_ON_OPENING: true, OrderTypeLIMIT_ON_CLOSE: true,
	OrderTypeMARKET_ON_CLOSE: true, OrderTypeLIMIT_IF_TOUCHED: true,
	OrderTypeMARKET_IF_TOUCHED: true,
}

func ParseOrderType(s string) (OrderType, error) {
	ot := OrderType(s)
	if !validOrderTypes[ot] {
		return "", fmt.Errorf("invalid order_type: %s", s)
	}
	return ot, nil
}

type OrderSide string

const (
	OrderSideBUY  OrderSide = "BUY"
	OrderSideSELL OrderSide = "SELL"
)

func ParseOrderSide(s string) (OrderSide, error) {
	switch OrderSide(s) {
	case OrderSideBUY, OrderSideSELL:
		return OrderSide(s), nil
	default:
		return "", fmt.Errorf("invalid order_side: %s", s)
	}
}

type TimeInForce string

const (
	TimeInForceDAY TimeInForce = "DAY"
	TimeInForceGTC TimeInForce = "GTC"
)

func ParseTimeInForce(s string) (TimeInForce, error) {
	switch TimeInForce(s) {
	case TimeInForceDAY, TimeInForceGTC:
		return TimeInForce(s), nil
	default:
		return "", fmt.Errorf("invalid time_in_force: %s", s)
	}
}

type OrderIdType string

const (
	OrderIdTypeCLIENT OrderIdType = "CLIENT"
	OrderIdTypeSNB    OrderIdType = "SNB"
)

type OrderStatus string

const (
	OrderStatusINVALID           OrderStatus = "INVALID"
	OrderStatusEXPIRED           OrderStatus = "EXPIRED"
	OrderStatusNO_REPORT         OrderStatus = "NO_REPORT"
	OrderStatusWAIT_REPORT       OrderStatus = "WAIT_REPORT"
	OrderStatusREPORTED          OrderStatus = "REPORTED"
	OrderStatusPART_CONCLUDED    OrderStatus = "PART_CONCLUDED"
	OrderStatusCONCLUDED         OrderStatus = "CONCLUDED"
	OrderStatusWITHDRAWING       OrderStatus = "WITHDRAWING"
	OrderStatusWAIT_WITHDRAW     OrderStatus = "WAIT_WITHDRAW"
	OrderStatusPART_WAIT_WITHDRAW OrderStatus = "PART_WAIT_WITHDRAW"
	OrderStatusPART_WITHDRAW     OrderStatus = "PART_WITHDRAW"
	OrderStatusWITHDRAWED        OrderStatus = "WITHDRAWED"
	OrderStatusREPLACING         OrderStatus = "REPLACING"
	OrderStatusWAIT_REPLACE      OrderStatus = "WAIT_REPLACE"
	OrderStatusREPLACED          OrderStatus = "REPLACED"
)

type TradingHours string

const (
	TradingHoursONLY_RTH      TradingHours = "ONLY_RTH"
	TradingHoursNOT_ONLY_RTH  TradingHours = "NOT_ONLY_RTH"
	TradingHoursOVERNIGHT     TradingHours = "OVERNIGHT"
	TradingHoursDAY_OVERNIGHT TradingHours = "DAY_OVERNIGHT"
)
