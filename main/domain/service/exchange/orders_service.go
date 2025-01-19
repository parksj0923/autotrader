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
	OrderChanceURL = "https://api.upbit.com/v1/orders/chance"
	OrderUrl       = "https://api.upbit.com/v1/order"
	OrdersUrl      = "https://api.upbit.com/v1/orders"
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
	for k, v := range params {
		qParams = append(qParams, resty.QueryParam{
			Key:   k,
			Value: v,
		})
	}

	resp, err := service.resty.MakeRequest(ctx, nil, header).Get(OrderChanceURL, qParams...)
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

// GetOrder : 단일 주문 내역을 조회한다.
// origin api : GET /order
// request의 uuidOrIdentifier와 isIdentifier를 통해 uuid로 조회할지 identifier로 조회할지 구분
func (service *exchangeService) GetOrder(ctx context.Context, uuidOrIdentifier string, isIdentifier bool) (*protocols.OrderResponse, error) {
	params := map[string]string{}
	if isIdentifier {
		params["identifier"] = uuidOrIdentifier
	} else {
		params["uuid"] = uuidOrIdentifier
	}

	token, err := utils.GenerateJWT(params)
	if err != nil {
		return nil, fmt.Errorf("JWT 생성 실패: %w", err)
	}

	header := map[string]string{
		"Authorization": "Bearer " + token,
	}

	var qParams []resty.QueryParam
	for k, v := range params {
		qParams = append(qParams, resty.QueryParam{Key: k, Value: v})
	}

	resp, err := service.resty.
		MakeRequest(ctx, nil, header).
		Get(OrderUrl, qParams...)

	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API 응답 오류: %d, %s", resp.StatusCode(), resp.String())
	}

	// 6) 응답 파싱
	var result protocols.OrderResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("JSON 파싱 실패: %w", err)
	}
	return &result, nil
}

// CreateOrder : 매수/매도 주문을 생성합니다
// origin api : POST /orders
func (service *exchangeService) CreateOrder(ctx context.Context, req protocols.CreateOrderRequest) (*protocols.OrderResponse, error) {
	params := map[string]string{
		"market":   req.Market,
		"side":     req.Side,
		"ord_type": req.OrdType,
	}
	if req.Price != "" {
		params["price"] = req.Price
	}
	if req.Volume != "" {
		params["volume"] = req.Volume
	}
	if req.Identifier != "" {
		params["identifier"] = req.Identifier
	}
	if req.TimeInForce != "" {
		params["time_in_force"] = req.TimeInForce
	}

	token, err := utils.GenerateJWT(params)
	if err != nil {
		return nil, fmt.Errorf("JWT 생성 실패: %w", err)
	}

	header := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	body := params
	resp, err := service.resty.
		MakeRequest(ctx, body, header).
		Post(OrdersUrl)
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

// CancelOrder : 단일 주문을 취소한다.
// origin api : DELETE /v1/order
func (service *exchangeService) CancelOrder(ctx context.Context, uuidOrIdentifier string, isIdentifier bool) (*protocols.OrderResponse, error) {
	params := map[string]string{}
	if isIdentifier {
		params["identifier"] = uuidOrIdentifier
	} else {
		params["uuid"] = uuidOrIdentifier
	}

	token, err := utils.GenerateJWT(params)
	if err != nil {
		return nil, fmt.Errorf("JWT 생성 실패: %w", err)
	}

	header := map[string]string{
		"Authorization": "Bearer " + token,
		"Accept":        "application/json",
	}

	var qParams []resty.QueryParam
	for k, v := range params {
		qParams = append(qParams, resty.QueryParam{Key: k, Value: v})
	}

	resp, err := service.resty.
		MakeRequest(ctx, nil, header).
		Delete(OrderUrl, qParams...)
	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API 응답 오류: %d, %s", resp.StatusCode(), resp.String())
	}

	var result protocols.OrderResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("JSON 파싱 실패: %w", err)
	}
	return &result, nil
}
