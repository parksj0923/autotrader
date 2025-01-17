package route

import (
	exchangeServ "autotrader/main/domain/service/exchange"
	"autotrader/main/handler"
	"autotrader/main/infra"
)

func ExchangeRoute() {
	restyClient := infra.Contexts.Resty
	exchangeService := exchangeServ.NewExchangeService(restyClient)
	exchangeHandler := handler.NewExchangeHandler(exchangeService)

	exchange := infra.Contexts.Router.Group("/exchange")
	v1 := exchange.Group("/v1")

	v1.Get("/accounts", exchangeHandler.GetAccountInfoHandler())
}