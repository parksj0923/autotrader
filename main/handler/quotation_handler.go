package handler

import (
	"autotrader/main/common/fiberhelper/response"
	"autotrader/main/domain/service/quotation"
	protocols "autotrader/main/protocols/quotation"
	"github.com/gofiber/fiber/v2"
)

func NewQuotationHandler(service quotation.QuotationService) quotationHandler {
	return quotationHandler{service}
}

type quotationHandler struct {
	service quotation.QuotationService
}

func (handler quotationHandler) GetMarkets() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userContext := c.UserContext()
		var param = new(protocols.Market)
		if err := c.QueryParser(param); err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		result, err := handler.service.GetMarkets(userContext, param.IsDetails)
		if err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		return response.Ext{Ctx: c}.Ok(result)
	}
}
