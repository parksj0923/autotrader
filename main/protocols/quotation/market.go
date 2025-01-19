package quotation

type Market struct {
	IsDetails bool `json:"isDetails"`
}

// MarketResponse : 종목 코드 조회 결과
type MarketResponse struct {
	Market      string `json:"market"`       // e.g. "KRW-BTC"
	KoreanName  string `json:"korean_name"`  // e.g. "비트코인"
	EnglishName string `json:"english_name"` // e.g. "Bitcoin"`
	MarketEvent struct {
		Warning bool `json:"warning"` // 유의 종목인지 여부
		Caution struct {
			PriceFluctuations            bool `json:"PRICE_FLUCTUATIONS"`
			TradingVolumeSoaring         bool `json:"TRADING_VOLUME_SOARING"`
			DepositAmountSoaring         bool `json:"DEPOSIT_AMOUNT_SOARING"`
			GlobalPriceDifferences       bool `json:"GLOBAL_PRICE_DIFFERENCES"`
			ConcentrationOfSmallAccounts bool `json:"CONCENTRATION_OF_SMALL_ACCOUNTS"`
		} `json:"caution"`
	} `json:"market_event"`
}
