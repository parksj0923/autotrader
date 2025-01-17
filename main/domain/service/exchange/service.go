package exchange

import (
	"autotrader/main/common/resty"
	"autotrader/main/protocols/exchange"
	"context"
)

type ExchangeService interface {
	GetAccounts(ctx context.Context) ([]exchange.AccountResponse, error)
}

type exchangeService struct {
	resty resty.RestyClient
}

func NewExchangeService(resty resty.RestyClient) ExchangeService {
	return &exchangeService{resty}
}
