package exchange

import (
	"autotrader/main/common/resty"
	"autotrader/main/domain/service/exchange"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetAccounts(t *testing.T) {
	ctx := context.Background()
	var restyClient = resty.NewDefaultRestyClient(true, 10*time.Second)
	assetService := exchange.NewExchangeService(restyClient)
	result, err := assetService.GetAccounts(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
