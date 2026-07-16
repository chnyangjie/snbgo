# snbgo

Go client library for the Snowball Securities (雪盈证券) trading API.

## Installation

```bash
go get github.com/chnyangjie/snbgo
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/chnyangjie/snbgo"
)

func main() {
    cfg, err := snbgo.LoadFromEnv()
    if err != nil {
        panic(err)
    }

    client := snbgo.NewClient(cfg)

    // Login
    resp, err := client.Login()
    if err != nil {
        panic(err)
    }
    if !resp.Succeed() {
        panic(resp.MsgString())
    }

    // Get balance
    balance, _ := client.GetBalance()
    fmt.Println(balance.DataString())
}
```

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `SNB_ACCOUNT` | Yes | - | Account ID |
| `SNB_KEY` | Yes | - | API secret key |
| `SNB_SERVER` | No | `sandbox.snbsecurities.com` | API host |
| `SNB_PORT` | No | `443` | API port |
| `SNB_SCHEMA` | No | `https` | `http` or `https` |
| `SNB_TIMEOUT` | No | `10000` | Request timeout (ms) |

## API Methods

- `Login()` — Authenticate and obtain access token
- `GetBalance()` — Query account balance
- `GetPositionList(securityType)` — Query current positions
- `GetOrderList(page, size, status, securityType)` — Query orders
- `GetOrderByID(orderID)` — Query single order
- `PlaceOrder(...)` — Place a trade order
- `CancelOrder(orderID, originOrderID)` — Cancel an order
- `GetTransactionList(page, size)` — Query transaction history
- `GetSecurityDetail(symbol)` — Query security details

## License

Apache License 2.0
