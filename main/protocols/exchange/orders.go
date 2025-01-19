package protocols

import "time"

type OrderChangeQParam struct {
	Market string `json:"market"`
}

type SingleOrderQParam struct {
	Uuid       string `json:"uuid"`
	Identifier string `json:"identifier"`
}

type OrderChanceResponse struct {
	Market struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		OrderTypes []string `json:"order_types"`
		OrderSides []string `json:"order_sides"`
		Bid        struct {
			Currency  string  `json:"currency"`
			PriceUnit *string `json:"price_unit"`
			MinTotal  string  `json:"min_total"`
		} `json:"bid"`
		Ask struct {
			Currency  string  `json:"currency"`
			PriceUnit *string `json:"price_unit"`
			MinTotal  string  `json:"min_total"`
		} `json:"ask"`
	} `json:"market"`

	BidAccount struct {
		Currency     string `json:"currency"`
		Balance      string `json:"balance"`
		Locked       string `json:"locked"`
		AvgBuyPrice  string `json:"avg_buy_price"`
		UnitCurrency string `json:"unit_currency"`
	} `json:"bid_account"`

	AskAccount struct {
		Currency     string `json:"currency"`
		Balance      string `json:"balance"`
		Locked       string `json:"locked"`
		AvgBuyPrice  string `json:"avg_buy_price"`
		UnitCurrency string `json:"unit_currency"`
	} `json:"ask_account"`
}

// OrderResponse : 주문 생성/취소/조회 시 반환되는 객체
type OrderResponse struct {
	UUID            string    `json:"uuid"`
	Side            string    `json:"side"`
	OrdType         string    `json:"ord_type"`
	Price           string    `json:"price"`
	State           string    `json:"state"`
	Market          string    `json:"market"`
	CreatedAt       time.Time `json:"created_at"`
	Volume          string    `json:"volume"`
	RemainingVolume string    `json:"remaining_volume"`
	ReservedFee     string    `json:"reserved_fee"`
	RemainingFee    string    `json:"remaining_fee"`
	PaidFee         string    `json:"paid_fee"`
	Locked          string    `json:"locked"`
	ExecutedVolume  string    `json:"executed_volume"`
	TradesCount     int       `json:"trades_count"`
	Trades          []Trade   `json:"trades"` // 체결 내역이 있을 경우
}

type Trade struct {
	Market    string    `json:"market"`
	UUID      string    `json:"uuid"`
	Price     string    `json:"price"`
	Volume    string    `json:"volume"`
	Funds     string    `json:"funds"`
	Side      string    `json:"side"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOrderRequest struct {
	Market      string `json:"market"`                  // e.g. "KRW-BTC"
	Side        string `json:"side"`                    // "bid" or "ask"
	Volume      string `json:"volume"`                  // (필수/옵션) 주문 수량
	Price       string `json:"price"`                   // (필수/옵션) 1코인당 주문 가격
	OrdType     string `json:"ord_type"`                // "limit", "price", "market", "best"
	Identifier  string `json:"identifier"`              // (옵션) 주문 식별 값
	TimeInForce string `json:"time_in_force,omitempty"` // "ioc", "fok" (옵션)
}
