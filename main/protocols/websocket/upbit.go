package websocket

// UpbitTickerMessage : 업비트 'ticker' 타입 메시지를 매핑하는 구조체
type UpbitTickerMessage struct {
	Type string `json:"type"` // e.g. "ticker"
	Code string `json:"code"` // e.g. "KRW-BTC"

	OpeningPrice     float64 `json:"opening_price"`      // 시가
	HighPrice        float64 `json:"high_price"`         // 고가
	LowPrice         float64 `json:"low_price"`          // 저가
	TradePrice       float64 `json:"trade_price"`        // 현재가
	PrevClosingPrice float64 `json:"prev_closing_price"` // 전일 종가

	AccTradePrice     float64 `json:"acc_trade_price"`     // 누적 거래대금 (UTC 0시 기준)
	Change            string  `json:"change"`              // RISE, FALL, EVEN
	ChangePrice       float64 `json:"change_price"`        // 부호 없는 전일 대비
	SignedChangePrice float64 `json:"signed_change_price"` // 전일 대비 (부호 포함)
	ChangeRate        float64 `json:"change_rate"`         // 부호 없는 등락율
	SignedChangeRate  float64 `json:"signed_change_rate"`  // 등락율 (부호 포함)

	AskBid         string  `json:"ask_bid"`          // 매수/매도 구분 (ASK, BID)
	TradeVolume    float64 `json:"trade_volume"`     // 가장 최근 거래량
	AccTradeVolume float64 `json:"acc_trade_volume"` // 누적 거래량 (UTC 0시 기준)
	TradeDate      string  `json:"trade_date"`       // 최근 거래 일자 (UTC, yyyyMMdd)
	TradeTime      string  `json:"trade_time"`       // 최근 거래 시각 (UTC, HHmmss)

	TradeTimestamp int64   `json:"trade_timestamp"` // 체결 타임스탬프 (milliseconds)
	AccAskVolume   float64 `json:"acc_ask_volume"`  // 누적 매도량
	AccBidVolume   float64 `json:"acc_bid_volume"`  // 누적 매수량

	Highest52WeekPrice float64 `json:"highest_52_week_price"` // 52주 최고가
	Highest52WeekDate  string  `json:"highest_52_week_date"`  // 52주 최고가 달성일 (yyyy-MM-dd)
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`  // 52주 최저가
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`   // 52주 최저가 달성일 (yyyy-MM-dd)

	// market_state: 거래 상태
	//  - PREVIEW: 입금지원
	//  - ACTIVE: 거래지원가능
	//  - DELISTED: 거래지원종료
	MarketState        string `json:"market_state"`
	IsTradingSuspended bool   `json:"is_trading_suspended"` // Deprecated 필드지만, 문서상 존재

	DelistingDate *string `json:"delisting_date"` // 거래지원 종료일 (nullable)

	// market_warning: 유의 종목 여부
	//  - NONE: 해당 없음
	//  - CAUTION: 투자유의
	MarketWarning string `json:"market_warning"`

	Timestamp int64 `json:"timestamp"` // 타임스탬프 (milliseconds)

	AccTradePrice24h  float64 `json:"acc_trade_price_24h"`  // 24시간 누적 거래대금
	AccTradeVolume24h float64 `json:"acc_trade_volume_24h"` // 24시간 누적 거래량

	StreamType string `json:"stream_type"` // SNAPSHOT or REALTIME
}

// UpbitTradeMessage : 업비트 'trade' 타입 메시지를 매핑하는 구조체
type UpbitTradeMessage struct {
	Type             string  `json:"type"`               // "trade"
	Code             string  `json:"code"`               // 마켓 코드, ex) "KRW-BTC"
	Timestamp        int64   `json:"timestamp"`          // 타임스탬프 (millisecond)
	TradeDate        string  `json:"trade_date"`         // 체결 일자 (yyyy-MM-dd)
	TradeTime        string  `json:"trade_time"`         // 체결 시각 (HH:mm:ss)
	TradeTimestamp   int64   `json:"trade_timestamp"`    // 체결 타임스탬프 (millisecond)
	TradePrice       float64 `json:"trade_price"`        // 체결 가격
	TradeVolume      float64 `json:"trade_volume"`       // 체결량
	AskBid           string  `json:"ask_bid"`            // "ASK" 또는 "BID"
	PrevClosingPrice float64 `json:"prev_closing_price"` // 전일 종가
	Change           string  `json:"change"`             // "RISE", "EVEN", "FALL"
	ChangePrice      float64 `json:"change_price"`       // 부호 없는 전일 대비 값
	SequentialID     int64   `json:"sequential_id"`      // 체결 번호 (Unique)
	BestAskPrice     float64 `json:"best_ask_price"`     // 최우선 매도 호가
	BestAskSize      float64 `json:"best_ask_size"`      // 최우선 매도 잔량
	BestBidPrice     float64 `json:"best_bid_price"`     // 최우선 매수 호가
	BestBidSize      float64 `json:"best_bid_size"`      // 최우선 매수 잔량
	StreamType       string  `json:"stream_type"`        // "SNAPSHOT" 또는 "REALTIME"
}

// OrderbookUnit : 개별 호가 단위 정보
type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"` // 매도 호가
	BidPrice float64 `json:"bid_price"` // 매수 호가
	AskSize  float64 `json:"ask_size"`  // 매도 잔량
	BidSize  float64 `json:"bid_size"`  // 매수 잔량
}

// UpbitOrderbookMessage : 업비트 호가(orderbook) 메시지 구조체
type UpbitOrderbookMessage struct {
	Type           string          `json:"type"`            // "orderbook"
	Code           string          `json:"code"`            // 마켓 코드 (ex. "KRW-BTC")
	Timestamp      int64           `json:"timestamp"`       // 타임스탬프 (millisecond)
	TotalAskSize   float64         `json:"total_ask_size"`  // 호가 매도 총 잔량
	TotalBidSize   float64         `json:"total_bid_size"`  // 호가 매수 총 잔량
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"` // 호가 단위 정보 리스트
	Level          float64         `json:"level"`           // 호가 모아보기 단위 (Optional)
}

// UpbitMyOrderMessage : 업비트의 myOrder 타입 웹소켓 응답 메시지 구조체
type UpbitMyOrderMessage struct {
	Type            string  `json:"type"`             // "myOrder"
	Code            string  `json:"code"`             // 마켓 코드, ex) "KRW-BTC"
	UUID            string  `json:"uuid"`             // 주문 고유 아이디
	AskBid          string  `json:"ask_bid"`          // "ASK" 또는 "BID"
	OrderType       string  `json:"order_type"`       // "limit", "price", "market", "best" 등
	State           string  `json:"state"`            // 주문 상태 ("wait", "trade", "done", "cancel" 등)
	TradeUUID       string  `json:"trade_uuid"`       // 체결 고유 아이디 (trade 상태일 때)
	Price           float64 `json:"price"`            // 주문 가격 또는 체결 가격
	AvgPrice        float64 `json:"avg_price"`        // 평균 체결 가격
	Volume          float64 `json:"volume"`           // 주문량 또는 체결량
	RemainingVolume float64 `json:"remaining_volume"` // 체결 후 남은 주문 양
	ExecutedVolume  float64 `json:"executed_volume"`  // 체결된 양
	TradesCount     int     `json:"trades_count"`     // 해당 주문에 걸린 체결 수
	ReservedFee     float64 `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    float64 `json:"remaining_fee"`    // 남은 수수료
	PaidFee         float64 `json:"paid_fee"`         // 사용된 수수료
	Locked          float64 `json:"locked"`           // 거래에 사용중인 비용
	ExecutedFunds   float64 `json:"executed_funds"`   // 체결된 금액
	TimeInForce     string  `json:"time_in_force"`    // IOC, FOK 설정 (없을 수 있음)
	TradeFee        float64 `json:"trade_fee"`        // 체결 시 발생한 수수료 (trade 타입일 때만)
	IsMaker         *bool   `json:"is_maker"`         // 체결이 발생한 주문의 maker/taker 여부 (true/false, 없을 수 있음)
	Identifier      string  `json:"identifier"`       // 조회용 사용자 지정값
	TradeTimestamp  int64   `json:"trade_timestamp"`  // 체결 타임스탬프 (millisecond)
	OrderTimestamp  int64   `json:"order_timestamp"`  // 주문 타임스탬프 (millisecond)
	Timestamp       int64   `json:"timestamp"`        // 타임스탬프 (millisecond)
	StreamType      string  `json:"stream_type"`      // 스트림 타입, 예: "REALTIME"
}

// Asset : 개별 자산 정보
type Asset struct {
	Currency string  `json:"currency"` // 화폐 코드, ex) "KRW"
	Balance  float64 `json:"balance"`  // 주문가능 수량
	Locked   float64 `json:"locked"`   // 주문 중 묶여있는 수량
}

// UpbitMyAssetMessage : 업비트 myAsset 타입 응답 메시지 구조체
type UpbitMyAssetMessage struct {
	Type           string  `json:"type"`            // "myAsset"
	AssetUUID      string  `json:"asset_uuid"`      // 자산 고유 아이디
	Assets         []Asset `json:"assets"`          // 자산 리스트
	AssetTimestamp int64   `json:"asset_timestamp"` // 자산 타임스탬프 (millisecond)
	Timestamp      int64   `json:"timestamp"`       // 타임스탬프 (millisecond)
	StreamType     string  `json:"stream_type"`     // 스트림 타입 (예: "REALTIME")
}
