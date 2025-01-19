package handler

import (
	fiberhelpers "autotrader/main/common/fiberhelper"
	"autotrader/main/common/fiberhelper/response"
	"autotrader/main/domain/service/exchange"
	protocols "autotrader/main/protocols/exchange"
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

func (handler exchangeHandler) GetOrderChanceInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userContext := c.UserContext()
		var param = new(protocols.OrderChangeQParam)
		if err := c.QueryParser(param); err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		result, err := handler.service.GetOrderChance(userContext, param.Market)
		if err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		return response.Ext{Ctx: c}.Ok(result)
	}
}

func (handler exchangeHandler) GetSingleOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userContext := c.UserContext()
		var param = new(protocols.SingleOrderQParam)
		if err := c.QueryParser(param); err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		uuid := param.Uuid
		identifier := param.Identifier
		uuidOrIdentifier := ""
		isIdentifier := false
		if identifier != "" {
			uuidOrIdentifier = identifier
			isIdentifier = true
		} else {
			uuidOrIdentifier = uuid
		}
		result, err := handler.service.GetOrder(userContext, uuidOrIdentifier, isIdentifier)
		if err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		return response.Ext{Ctx: c}.Ok(result)
	}
}

func (handler exchangeHandler) CreateOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userContext := c.UserContext()
		request := fiberhelpers.RequestParse[protocols.CreateOrderRequest](c)
		result, err := handler.service.CreateOrder(userContext, request)
		if err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		return response.Ext{Ctx: c}.Ok(result)
	}
}

func (handler exchangeHandler) CancelOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userContext := c.UserContext()
		var param = new(protocols.SingleOrderQParam)
		if err := c.QueryParser(param); err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		uuid := param.Uuid
		identifier := param.Identifier
		uuidOrIdentifier := ""
		isIdentifier := false
		if identifier != "" {
			uuidOrIdentifier = identifier
			isIdentifier = true
		} else {
			uuidOrIdentifier = uuid
		}
		result, err := handler.service.CancelOrder(userContext, uuidOrIdentifier, isIdentifier)
		if err != nil {
			return response.Ext{Ctx: c}.Error(err)
		}
		return response.Ext{Ctx: c}.Ok(result)

	}
}
