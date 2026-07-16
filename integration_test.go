package snbgo

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func setupClient(t *testing.T) *Client {
	t.Helper()
	account := os.Getenv("SNB_ACCOUNT")
	key := os.Getenv("SNB_KEY")
	if account == "" || key == "" {
		t.Skip("SNB_ACCOUNT and SNB_KEY not set")
	}
	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv: %v", err)
	}
	return NewClient(cfg)
}

func printJSON(t *testing.T, label string, resp *APIResponse) {
	t.Helper()
	if resp == nil {
		return
	}
	var raw interface{}
	if err := json.Unmarshal(resp.ResultData, &raw); err == nil {
		b, _ := json.MarshalIndent(raw, "", "  ")
		t.Logf("%s:\n%s", label, string(b))
	} else {
		t.Logf("%s: %s", label, resp.DataString())
	}
}

func TestLogin(t *testing.T) {
	client := setupClient(t)

	resp, err := client.Login()
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("Login failed: %s - %s", resp.ResultCode, resp.Message())
	}

	var data TokenData
	if err := resp.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData: %v", err)
	}
	if data.AccessToken == "" {
		t.Fatal("access_token is empty")
	}
	if data.ExpiryTime == 0 {
		t.Fatal("expiry_time is 0")
	}
	t.Logf("Login OK: token=%s..., expiry=%d", data.AccessToken[:8], data.ExpiryTime)
}

func TestGetBalance(t *testing.T) {
	client := setupClient(t)

	resp, err := client.GetBalance()
	if err != nil {
		t.Fatalf("GetBalance failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("GetBalance failed: %s - %s", resp.ResultCode, resp.Message())
	}

	var data BalanceData
	if err := resp.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData: %v", err)
	}
	t.Logf("Balance: cash=%.2f %s, net_liquidation=%.2f, leverage=%.2f",
		data.Cash, data.Currency, data.NetLiquidationValue, data.Leverage)
	for _, item := range data.BalanceDetailItems {
		t.Logf("  %s: cash=%.2f", item.Currency, item.Cash)
	}
}

func TestGetPositionList(t *testing.T) {
	client := setupClient(t)

	resp, err := client.GetPositionList("STK,OPT,WAR,IOPT,FUT")
	if err != nil {
		t.Fatalf("GetPositionList failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("GetPositionList failed: %s - %s", resp.ResultCode, resp.Message())
	}

	var positions []PositionData
	if err := resp.UnmarshalData(&positions); err != nil {
		t.Fatalf("UnmarshalData: %v", err)
	}
	t.Logf("Positions: %d items", len(positions))
	for i, p := range positions {
		if i >= 5 {
			t.Logf("  ... and %d more", len(positions)-5)
			break
		}
		t.Logf("  %s %s: qty=%.0f, avg=%.2f, market=%.2f, pnl=%.2f",
			p.Symbol, p.SecurityType, p.Position, p.AveragePrice, p.MarketPrice, p.RealizedPnL)
	}
}

func TestGetOrderList(t *testing.T) {
	client := setupClient(t)

	resp, err := client.GetOrderList(1, 10, "ALL", "STK,OPT,WAR,IOPT,FUT")
	if err != nil {
		t.Fatalf("GetOrderList failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("GetOrderList failed: %s - %s", resp.ResultCode, resp.Message())
	}

	var data OrderListData
	if err := resp.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData: %v", err)
	}
	t.Logf("Orders: count=%d, page=%d, size=%d", data.Count, data.Page, data.Size)
	for i, o := range data.Items {
		if i >= 5 {
			t.Logf("  ... and %d more", len(data.Items)-5)
			break
		}
		t.Logf("  %s %s %s: qty=%d, price=%.2f, status=%s",
			o.Symbol, o.Side, o.SecurityType, o.Quantity, o.Price, o.Status)
	}
}

func TestGetOrderByID(t *testing.T) {
	client := setupClient(t)

	// First get an order ID from the order list
	listResp, err := client.GetOrderList(1, 1, "ALL", "STK")
	if err != nil {
		t.Fatalf("GetOrderList failed: %v", err)
	}
	if !listResp.Succeed() {
		t.Fatalf("GetOrderList failed: %s - %s", listResp.ResultCode, listResp.Message())
	}

	var listData OrderListData
	if err := listResp.UnmarshalData(&listData); err != nil {
		t.Fatalf("UnmarshalData: %v", err)
	}
	if len(listData.Items) == 0 {
		t.Skip("No orders to query")
	}

	orderID := listData.Items[0].ID
	if orderID == "" {
		t.Skip("Order ID is empty")
	}

	resp, err := client.GetOrderByID(orderID)
	if err != nil {
		t.Fatalf("GetOrderByID failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("GetOrderByID failed: %s - %s", resp.ResultCode, resp.Message())
	}

	printJSON(t, "OrderByID", resp)
}

func TestGetTransactionList(t *testing.T) {
	client := setupClient(t)

	resp, err := client.GetTransactionList(1, 10, "", 0, 0)
	if err != nil {
		t.Fatalf("GetTransactionList failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("GetTransactionList failed: %s - %s", resp.ResultCode, resp.Message())
	}

	printJSON(t, "Transactions", resp)

	var data TransactionListData
	if err := resp.UnmarshalData(&data); err != nil {
		t.Fatalf("UnmarshalData: %v", err)
	}
	t.Logf("Transactions: count=%d, page=%d, size=%d", data.Count, data.Page, data.Size)
	for i, tx := range data.Items {
		if i >= 5 {
			t.Logf("  ... and %d more", len(data.Items)-5)
			break
		}
		t.Logf("  %s %s %s: qty=%.0f, price=%.2f, status=%s",
			tx.Symbol, tx.Side, tx.SecurityType, tx.Quantity, tx.Price, tx.Status)
	}
}

func TestGetSecurityDetail(t *testing.T) {
	client := setupClient(t)

	resp, err := client.GetSecurityDetail("AAPL")
	if err != nil {
		t.Fatalf("GetSecurityDetail failed: %v", err)
	}
	if !resp.Succeed() {
		t.Fatalf("GetSecurityDetail failed: %s - %s", resp.ResultCode, resp.Message())
	}

	printJSON(t, "SecurityDetail", resp)
}

func TestIntegration(t *testing.T) {
	client := setupClient(t)

	// 1. Login
	loginResp, err := client.Login()
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	if !loginResp.Succeed() {
		t.Fatalf("Login failed: %s - %s", loginResp.ResultCode, loginResp.Message())
	}
	fmt.Println("✓ Login OK")

	// 2. Get Balance
	balanceResp, err := client.GetBalance()
	if err != nil {
		t.Fatalf("GetBalance: %v", err)
	}
	if !balanceResp.Succeed() {
		t.Fatalf("GetBalance failed: %s - %s", balanceResp.ResultCode, balanceResp.Message())
	}
	printJSON(t, "Balance", balanceResp)
	fmt.Println("✓ GetBalance OK")

	// 3. Get Positions
	positionResp, err := client.GetPositionList("STK,OPT,WAR,IOPT,FUT")
	if err != nil {
		t.Fatalf("GetPositionList: %v", err)
	}
	if !positionResp.Succeed() {
		t.Fatalf("GetPositionList failed: %s - %s", positionResp.ResultCode, positionResp.Message())
	}
	printJSON(t, "Positions", positionResp)
	fmt.Println("✓ GetPositionList OK")

	// 4. Get Orders
	orderResp, err := client.GetOrderList(1, 10, "ALL", "STK,OPT,WAR,IOPT,FUT")
	if err != nil {
		t.Fatalf("GetOrderList: %v", err)
	}
	if !orderResp.Succeed() {
		t.Fatalf("GetOrderList failed: %s - %s", orderResp.ResultCode, orderResp.Message())
	}
	printJSON(t, "Orders", orderResp)
	fmt.Println("✓ GetOrderList OK")

	// 5. Get Transactions
	transactionResp, err := client.GetTransactionList(1, 10, "", 0, 0)
	if err != nil {
		t.Fatalf("GetTransactionList: %v", err)
	}
	if !transactionResp.Succeed() {
		t.Fatalf("GetTransactionList failed: %s - %s", transactionResp.ResultCode, transactionResp.Message())
	}
	printJSON(t, "Transactions", transactionResp)
	fmt.Println("✓ GetTransactionList OK")

	// 6. Get Security Detail
	securityResp, err := client.GetSecurityDetail("AAPL")
	if err != nil {
		t.Fatalf("GetSecurityDetail: %v", err)
	}
	if !securityResp.Succeed() {
		t.Fatalf("GetSecurityDetail failed: %s - %s", securityResp.ResultCode, securityResp.Message())
	}
	printJSON(t, "SecurityDetail", securityResp)
	fmt.Println("✓ GetSecurityDetail OK")
}
