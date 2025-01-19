package quotation

import (
	"autotrader/main/common/resty"
	protocols "autotrader/main/protocols/quotation"
	"context"
)

type QuotationService interface {
	GetMarkets(ctx context.Context, isDetails bool) ([]protocols.MarketResponse, error)
}
type quotationService struct {
	resty resty.RestyClient
}

func NewQuotationService(resty resty.RestyClient) QuotationService {
	return &quotationService{resty}
}
