package exchange

import (
	"autotrader/main/common/resty"
	"autotrader/main/protocols/exchange"
	"autotrader/main/utils"
	"context"
	"encoding/json"
	"fmt"
)

const (
	ORDERCHANGEURL = "https://api.upbit.com/v1/orders/chance"
	ORDERSURL      = "https://api.upbit.com/v1/orders"
)

// GetOrderChance : 특정 market에 대한 주문 가능 정보(최소 주문 금액, 잔고 등)를 조회합니다
// origin api : GET /orders/chance
func (service *exchangeService) GetOrderChance(ctx context.Context, market string) (*protocols.OrderChanceResponse, error) {
	params := map[string]string{
		"market": market,
	}

	token, err := utils.GenerateJWT(params)
	if err != nil {
		return nil, fmt.Errorf("JWT 생성 실패: %w", err)
	}
	header := map[string]string{
		"Authorization": "Bearer " + token,
	}

	qParams := make([]resty.QueryParam, 0)
	for key, value := range params {
		qParams = append(qParams, resty.QueryParam{
			Key:   key,
			Value: value,
		})
	}

	resp, err := service.resty.MakeRequest(ctx, nil, header).Get(ORDERCHANGEURL, qParams...)
	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API 응답 오류: %d, %s", resp.StatusCode(), resp.String())
	}

	var result protocols.OrderChanceResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("JSON 파싱 실패: %w", err)
	}
	return &result, nil
}

// CreateOrder : 매수/매도 주문을 생성합니다
// origin api : POST /orders
// market, side, volume, price, ord_type가 필수이며, identifier는 옵션
func (service *exchangeService) CreateOrder(ctx context.Context, market, side, volume, price, ordType,
	identifier string) (*protocols.OrderResponse, error) {

	params := map[string]string{
		"market":   market,  // e.g. "KRW-BTC"
		"side":     side,    // "bid"(매수) or "ask"(매도)
		"volume":   volume,  // 수량 (ord_type=market+매도일 때 필수)
		"price":    price,   // 가격 (ord_type=limit or market+매수 시 필수)
		"ord_type": ordType, // "limit", "price", "market"
	}
	// identifier는 옵션
	if identifier != "" {
		params["identifier"] = identifier
	}

	token, err := utils.GenerateJWT(params)
	if err != nil {
		return nil, fmt.Errorf("JWT 생성 실패: %w", err)
	}

	header := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := service.resty.MakeRequest(ctx, nil, header).Post(ORDERSURL)
	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("API 응답 오류: %d, %s", resp.StatusCode(), resp.String())
	}

	var result protocols.OrderResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("JSON 파싱 실패: %w", err)
	}
	return &result, nil
}
