package quotation

import (
	"autotrader/main/common/resty"
	"autotrader/main/domain/service/quotation"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetOrderChance(t *testing.T) {
	ctx := context.Background()
	var restyClient = resty.NewDefaultRestyClient(true, 10*time.Second)
	marketService := quotation.NewQuotationService(restyClient)
	result, err := marketService.GetMarkets(ctx, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
