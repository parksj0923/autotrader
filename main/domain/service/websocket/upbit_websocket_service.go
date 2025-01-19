package websocket

import (
	protocols "autotrader/main/protocols/websocket"
	upbitWs "autotrader/main/protocols/websocket"
	"autotrader/main/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
)

func (s *websocketService) Start(ctx context.Context, subs []Subscription) error {
	conn, _, err := websocket.DefaultDialer.Dial(s.url, nil)
	if err != nil {
		return err
	}
	s.wsConn = conn
	log.Println("[UpbitWS] 연결 성공")

	subMsg := buildSubscribeMessage(utils.RandomWsUuid, subs)
	if err := conn.WriteJSON(subMsg); err != nil {
		return err
	}

	go s.readLoop()

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		select {
		case <-interrupt:
			log.Println("[UpbitWS] 인터럽트 시그널 수신, Stop 호출")
			s.Stop()
		case <-ctx.Done():
			log.Println("[UpbitWS] Context 종료, Stop 호출")
			s.Stop()
		}
	}()

	return nil
}

func (s *websocketService) Stop() error {
	close(s.doneChan)
	if s.wsConn != nil {
		s.wsConn.Close()
	}
	log.Println("[UpbitWS] 종료")
	return nil
}

func (s *websocketService) readLoop() {
	defer func() { log.Println("[UpbitWS] readLoop 종료") }()
	for {
		select {
		case <-s.doneChan:
			return
		default:
			_, msg, err := s.wsConn.ReadMessage()
			if err != nil {
				log.Printf("[UpbitWS] ReadMessage 에러: %v\n", err)
				return
			}
			parsed, err := parseUpbitMessage(msg)
			if err != nil {
				log.Printf("[UpbitWS] parseUpbitMessage 에러: %v\n", err)
				continue
			}
			switch val := parsed.(type) {
			case upbitWs.UpbitTickerMessage:
				log.Printf("[Ticker] %s 현재가: %.2f", val.Code, val.TradePrice)
			case upbitWs.UpbitTradeMessage:
				log.Printf("[Trade] %s 체결가: %.2f, 체결량: %.4f", val.Code, val.TradePrice, val.TradeVolume)
			case upbitWs.UpbitOrderbookMessage:
				if len(val.OrderbookUnits) > 0 {
					top := val.OrderbookUnits[0]
					log.Printf("[Orderbook] %s 매도: %.2f(%.4f), 매수: %.2f(%.4f)",
						val.Code, top.AskPrice, top.AskSize, top.BidPrice, top.BidSize)
				}
			case upbitWs.UpbitMyOrderMessage:
				log.Printf("[MyOrder] 주문ID=%s, 상태=%s", val.UUID, val.State)
			case upbitWs.UpbitMyAssetMessage:
				log.Printf("[MyAsset] 자산UUID=%s, 자산 개수=%d", val.AssetUUID, len(val.Assets))
			default:
				log.Println("[UpbitWS] 알 수 없는 메시지 타입")
			}
		}
	}
}

func buildSubscribeMessage(ticket string, subs []Subscription) []interface{} {
	msg := []interface{}{
		map[string]string{"ticket": ticket},
	}

	for _, sub := range subs {
		msg = append(msg, sub.ToMap())
	}

	msg = append(msg, map[string]string{"format": "DEFAULT"})
	return msg
}

func parseUpbitMessage(raw []byte) (interface{}, error) {
	var wrap struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &wrap); err != nil {
		return nil, err
	}

	switch wrap.Type {
	case "ticker":
		var tm protocols.UpbitTickerMessage
		if err := json.Unmarshal(raw, &tm); err != nil {
			return nil, err
		}
		return tm, nil
	case "trade":
		var tr protocols.UpbitTradeMessage
		if err := json.Unmarshal(raw, &tr); err != nil {
			return nil, err
		}
		return tr, nil
	case "orderbook":
		var ob protocols.UpbitOrderbookMessage
		if err := json.Unmarshal(raw, &ob); err != nil {
			return nil, err
		}
		return ob, nil
	case "myOrder":
		var mo protocols.UpbitMyOrderMessage
		if err := json.Unmarshal(raw, &mo); err != nil {
			return nil, err
		}
		return mo, nil
	case "myAsset":
		var ma protocols.UpbitMyAssetMessage
		if err := json.Unmarshal(raw, &ma); err != nil {
			return nil, err
		}
		return ma, nil
	default:
		return nil, fmt.Errorf("unknown message type: %s", wrap.Type)
	}
}

type Subscription interface {
	ToMap() map[string]interface{}
}

// TicketField : WebSocket 요청의 첫 번째 객체, 세션 식별용 티켓 필드
type TicketField struct {
	Ticket string `json:"ticket"`
}

// TickerTypeField : 현재가(ticker) 정보를 구독하기 위한 요청 객체
type TickerTypeField struct {
	Type           string   `json:"type"`                       // "ticker"
	Codes          []string `json:"codes"`                      // 요청할 마켓 코드 리스트 (대문자 필요)
	IsOnlySnapshot bool     `json:"is_only_snapshot,omitempty"` // 옵션: 스냅샷만 수신
	IsOnlyRealtime bool     `json:"is_only_realtime,omitempty"` // 옵션: 실시간 데이터만 수신
}

func (t TickerTypeField) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"type":  t.Type,
		"codes": t.Codes,
	}
	if t.IsOnlySnapshot {
		m["is_only_snapshot"] = t.IsOnlySnapshot
	}
	if t.IsOnlyRealtime {
		m["is_only_realtime"] = t.IsOnlyRealtime
	}
	return m
}

// TradeTypeField : 체결(trade) 정보를 구독하기 위한 요청 객체
type TradeTypeField struct {
	Type           string   `json:"type"`                       // "trade"
	Codes          []string `json:"codes"`                      // 대문자로 요청해야 함
	IsOnlySnapshot bool     `json:"is_only_snapshot,omitempty"` // 생략 시 false
	IsOnlyRealtime bool     `json:"is_only_realtime,omitempty"` // 생략 시 false
}

func (t TradeTypeField) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"type":  t.Type,
		"codes": t.Codes,
	}
	if t.IsOnlySnapshot {
		m["is_only_snapshot"] = t.IsOnlySnapshot
	}
	if t.IsOnlyRealtime {
		m["is_only_realtime"] = t.IsOnlyRealtime
	}
	return m
}

// OrderbookTypeField : 호가(orderbook) 정보를 구독하기 위한 요청 객체
type OrderbookTypeField struct {
	Type           string   `json:"type"`                       // "orderbook"
	Codes          []string `json:"codes"`                      // 대문자로 요청해야 함, 예: ["KRW-BTC"]
	Level          float64  `json:"level,omitempty"`            // 호가 모아보기 단위 (옵션)
	IsOnlySnapshot bool     `json:"is_only_snapshot,omitempty"` // 옵션, 기본 false
	IsOnlyRealtime bool     `json:"is_only_realtime,omitempty"` // 옵션, 기본 false
}

func (o OrderbookTypeField) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"type":  o.Type,
		"codes": o.Codes,
	}
	if o.Level != 0 {
		m["level"] = o.Level
	}
	if o.IsOnlySnapshot {
		m["is_only_snapshot"] = o.IsOnlySnapshot
	}
	if o.IsOnlyRealtime {
		m["is_only_realtime"] = o.IsOnlyRealtime
	}
	return m
}

// MyOrderTypeField : 내 주문(myOrder) 정보를 구독하기 위한 요청 객체
type MyOrderTypeField struct {
	Type  string   `json:"type"`            // "myOrder" 고정
	Codes []string `json:"codes,omitempty"` // 특정 마켓 코드 리스트. 생략 시 모든 마켓
}

func (m MyOrderTypeField) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"type": m.Type,
	}
	if len(m.Codes) > 0 {
		result["codes"] = m.Codes
	}
	return result
}

// MyAssetTypeField : 내 자산(myAsset) 정보를 구독하기 위한 요청 객체
type MyAssetTypeField struct {
	Type string `json:"type"` // "myAsset" 고정
}

func (m MyAssetTypeField) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type": m.Type,
	}
}

// FormatField : 응답 포맷 지정 (예: SIMPLE, DEFAULT 등)
type FormatField struct {
	Format string `json:"format"`
}
