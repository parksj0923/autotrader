package infra

import (
	"autotrader/main/common/resty"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type globalApplicationContext struct {
	Router fiber.Router
	Resty  resty.RestyClient
}

func Init(router fiber.Router, resty resty.RestyClient) error {
	if Contexts.Router == nil {
		Contexts.Router = router
	} else {
		return errors.New("don't reallocate fiber's app")
	}
	Contexts.Resty = resty
	return nil
}

var Contexts = &globalApplicationContext{}
