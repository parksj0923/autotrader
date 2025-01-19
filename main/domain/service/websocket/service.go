package websocket

import (
	"context"
	"github.com/gorilla/websocket"
)

type WebsocketService interface {
	Start(ctx context.Context, subscriptions []Subscription) error
	Stop() error
}

type websocketService struct {
	wsConn   *websocket.Conn
	url      string
	doneChan chan struct{}
}

func NewUpbitWebsocketService(url string) WebsocketService {
	return &websocketService{
		url:      url,
		doneChan: make(chan struct{}),
	}
}
