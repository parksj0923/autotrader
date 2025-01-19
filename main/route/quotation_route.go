package route

import (
	quotationServ "autotrader/main/domain/service/quotation"
	"autotrader/main/handler"
	"autotrader/main/infra"
)

func QuotationRoute() {
	restyClient := infra.Contexts.Resty
	quotationService := quotationServ.NewQuotationService(restyClient)
	quotationHandler := handler.NewQuotationHandler(quotationService)

	quotation := infra.Contexts.Router.Group("/quotation")
	v1 := quotation.Group("/v1")

	v1.Get("/market", quotationHandler.GetMarkets())
}
