package protocols

import "time"

// ------------------ 주문 가능 정보 (GET /orders/chance) ------------------ //

type OrderChangeQParam struct {
	Market string `json:"market"`
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

// ------------------ 주문 생성/취소/조회 시 공통 (POST/DELETE/GET /orders) ------------------ //

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

// OrderListResponse: 주문 목록 조회( GET /orders ) 시 여러 건 반환
// - API 응답은 배열이므로, []OrderResponse 형태로 받으면 충분.
//   별도 구조체 없이 []OrderResponse로 직렬화하는 방식도 가능.
