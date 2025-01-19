package exchange

import (
	"autotrader/main/common/resty"
	"autotrader/main/domain/service/exchange"
	protocols "autotrader/main/protocols/exchange"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetOrderChance(t *testing.T) {
	ctx := context.Background()
	var restyClient = resty.NewDefaultRestyClient(true, 10*time.Second)
	assetService := exchange.NewExchangeService(restyClient)
	result, err := assetService.GetOrderChance(ctx, "KRW-BTC")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func TestPostOrder(t *testing.T) {
	ctx := context.Background()

	orderRequest := protocols.CreateOrderRequest{
		Market:      "KRW-DOGE",
		Side:        "",
		Volume:      "",
		Price:       "",
		OrdType:     "",
		Identifier:  "",
		TimeInForce: "",
	}
	var restyClient = resty.NewDefaultRestyClient(true, 10*time.Second)
	assetService := exchange.NewExchangeService(restyClient)
	result, err := assetService.CreateOrder(ctx, "KRW-BTC")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
