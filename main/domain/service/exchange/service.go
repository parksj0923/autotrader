package exchange

import (
	"autotrader/main/common/resty"
	"autotrader/main/protocols/exchange"
	"context"
)

type ExchangeService interface {
	GetAccounts(ctx context.Context) ([]protocols.AccountResponse, error)
	GetOrderChance(ctx context.Context, market string) (*protocols.OrderChanceResponse, error)
	CreateOrder(ctx context.Context, market, side, volume, price, ordType, identifier string) (*protocols.OrderResponse, error)
}

type exchangeService struct {
	resty resty.RestyClient
}

func NewExchangeService(resty resty.RestyClient) ExchangeService {
	return &exchangeService{resty}
}
