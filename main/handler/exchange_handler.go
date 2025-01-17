package handler

import (
	"autotrader/main/common/fiberhelper/response"
	"autotrader/main/domain/service/exchange"
	"github.com/gofiber/fiber/v2"
)

func NewExchangeHandler(service exchange.ExchangeService) exchangeHandler {
	return exchangeHandler{service}
}

type exchangeHandler struct {
	service exchange.ExchangeService
}

func (handler exchangeHandler) GetAccountInfoHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userContext := c.UserContext()
		result, err := handler.service.GetAccounts(userContext)
		if err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		return response.Ext{Ctx: c}.Ok(result)
	}
}
