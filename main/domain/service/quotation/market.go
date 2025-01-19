package quotation

import (
	"autotrader/main/common/resty"
	protocols "autotrader/main/protocols/quotation"
	"context"
	"encoding/json"
	"fmt"
)

const (
	MarketAllURL = "https://api.upbit.com/v1/market/all"
)

func (service *quotationService) GetMarkets(ctx context.Context, isDetails bool) ([]protocols.MarketResponse, error) {
	var qParams []resty.QueryParam
	if isDetails {
		qParams = append(qParams, resty.QueryParam{Key: "is_details", Value: true})
	}

	header := map[string]string{
		"Accept": "application/json",
	}

	resp, err := service.resty.
		MakeRequest(ctx, nil, header).
		Get(MarketAllURL, qParams...)

	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API 응답 오류: %d, %s", resp.StatusCode(), resp.String())
	}

	var result []protocols.MarketResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("JSON 파싱 실패: %w", err)
	}
	return result, nil
}
