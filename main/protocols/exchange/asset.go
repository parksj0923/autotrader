package protocols

// AccountResponse : 업비트 계좌 정보 구조체
type AccountResponse struct {
	Currency     string `json:"currency"`      // 화폐 단위 (e.g., KRW, BTC)
	Balance      string `json:"balance"`       // 보유 잔고
	Locked       string `json:"locked"`        // 주문 중 묶여있는 잔고
	AvgBuyPrice  string `json:"avg_buy_price"` // 매수 평균가
	UnitCurrency string `json:"unit_currency"` // 기준 화폐 (e.g., KRW)
}
