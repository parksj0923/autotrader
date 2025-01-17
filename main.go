package main

import (
	fiberhelpers "autotrader/main/common/fiberhelper"
	"autotrader/main/common/fiberhelper/middleware"
	"autotrader/main/common/resty"
	"autotrader/main/infra"
	"autotrader/main/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New(
		fiber.Config{ErrorHandler: fiberhelpers.DefaultErrorHandler},
	)
	app.Use(middleware.LogMiddleware())
	app.Use(fiberhelpers.NewRecover())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // 허용할 도메인
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	restyClient := resty.NewDefaultRestyClient(true, 10*time.Second)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	api := app.Group("/api")
	err := infra.Init(api, restyClient)
	if err != nil {
		log.Panic(err.Error())
		panic(err.Error())
	}

	route.ExchangeRoute()

	fiberhelpers.ListenWithGraceFullyShutdown(app, port)
}
