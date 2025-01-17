package upbit

import (
	"autotrader/main/common/resty"
	"autotrader/main/domain/service/exchange/asset"
	"context"
	"fmt"
	"testing"
	"time"
)

var restyClient = resty.NewDefaultRestyClient(true, 10*time.Second)

func TestGetAccounts(t *testing.T) {
	ctx := context.Background()
	assetService := asset.NewAssetService(restyClient)
	result, err := assetService.GetAccounts(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
