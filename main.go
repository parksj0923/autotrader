package main

import (
	fiberhelpers "autotrader/main/common/fiberhelper"
	"autotrader/main/common/fiberhelper/middleware"
	"autotrader/main/common/resty"
	"autotrader/main/domain/service/websocket"
	"autotrader/main/infra"
	"autotrader/main/route"
	"autotrader/main/utils"
	"context"
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
	route.QuotationRoute()

	upbitPublicWs := websocket.NewUpbitWebsocketService(utils.UpbitPublicWsUrl)
	subs := []websocket.Subscription{
		websocket.TickerTypeField{
			Type:           utils.Ticker,
			Codes:          []string{"KRW-DOGE", "KRW-BTC", "KRW-ETH"},
			IsOnlySnapshot: false,
			IsOnlyRealtime: true,
		},
	}

	wsCtx := context.Background()
	go func() {
		if err := upbitPublicWs.Start(wsCtx, subs); err != nil {
			log.Println("WebSocket 서비스 시작 오류:", err)
		}
	}()

	fiberhelpers.ListenWithGraceFullyShutdown(app, port)
}
