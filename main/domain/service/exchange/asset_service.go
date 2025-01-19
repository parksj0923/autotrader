package exchange

import (
	"autotrader/main/protocols/exchange"
	"autotrader/main/utils"
	"context"
	"encoding/json"
	"fmt"
)

const (
	AccountURL = "https://api.upbit.com/v1/accounts"
)

// GetAccounts : 업비트 전체 계좌 조회 API 호출
func (service *exchangeService) GetAccounts(ctx context.Context) ([]protocols.AccountResponse, error) {
	token, err := utils.GenerateJWT(nil)
	if err != nil {
		return nil, fmt.Errorf("JWT 생성 실패: %w", err)
	}

	header := map[string]string{
		"Authorization": "Bearer " + token,
	}
	resp, err := service.resty.MakeRequest(ctx, nil, header).Get(AccountURL)
	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API 응답 오류: %d, %s", resp.StatusCode(), resp.String())
	}

	var accounts []protocols.AccountResponse
	if err := json.Unmarshal(resp.Body(), &accounts); err != nil {
		return nil, fmt.Errorf("JSON 파싱 실패: %w", err)
	}

	return accounts, nil
}
