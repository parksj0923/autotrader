package exchange

import (
	"autotrader/main/common/resty"
	"autotrader/main/protocols/exchange"
	"context"
)

type ExchangeService interface {
	GetAccounts(ctx context.Context) ([]protocols.AccountResponse, error)
	GetOrderChance(ctx context.Context, market string) (*protocols.OrderChanceResponse, error)
	GetOrder(ctx context.Context, uuidOrIdentifier string, isIdentifier bool) (*protocols.OrderResponse, error)
	CreateOrder(ctx context.Context, req protocols.CreateOrderRequest) (*protocols.OrderResponse, error)
	CancelOrder(ctx context.Context, uuidOrIdentifier string, isIdentifier bool) (*protocols.OrderResponse, error)
}

type exchangeService struct {
	resty resty.RestyClient
}

func NewExchangeService(resty resty.RestyClient) ExchangeService {
	return &exchangeService{resty}
}
