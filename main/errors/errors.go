package errors

// 400 Bad Request 범주
const (
	ErrCreateAskError               = "create_ask_error"
	ErrCreateBidError               = "create_bid_error"
	ErrInsufficientFundsAsk         = "insufficient_funds_ask"
	ErrInsufficientFundsBid         = "insufficient_funds_bid"
	ErrUnderMinTotalAsk             = "under_min_total_ask"
	ErrUnderMinTotalBid             = "under_min_total_bid"
	ErrWithdrawAddressNotRegistered = "withdraw_address_not_registerd" // 문서 오타 그대로 사용
	ErrValidationError              = "validation_error"
)

// 401 Unauthorized 범주
const (
	ErrInvalidQueryPayload = "invalid_query_payload"
	ErrJWTVerification     = "jwt_verification"
	ErrExpiredAccessKey    = "expired_access_key"
	ErrNonceUsed           = "nonce_used"
	ErrNoAuthorizationIP   = "no_authorization_i_p"
	ErrOutOfScope          = "out_of_scope"
)

// ErrorMessages : 에러 코드와 기본 메시지 매핑
var ErrorMessages = map[string]string{
	// 400 Bad Request
	ErrCreateAskError:               "주문 요청 정보가 올바르지 않습니다. 파라미터 값이 올바른지 확인해주세요.",
	ErrCreateBidError:               "주문 요청 정보가 올바르지 않습니다. 파라미터 값이 올바른지 확인해주세요.",
	ErrInsufficientFundsAsk:         "매수/매도 가능 잔고가 부족합니다.",
	ErrInsufficientFundsBid:         "매수/매도 가능 잔고가 부족합니다.",
	ErrUnderMinTotalAsk:             "주문 요청 금액이 최소주문금액 미만입니다.",
	ErrUnderMinTotalBid:             "주문 요청 금액이 최소주문금액 미만입니다.",
	ErrWithdrawAddressNotRegistered: "허용되지 않은 출금 주소입니다. 허용 목록에 등록된 주소로만 출금이 가능합니다.",
	ErrValidationError:              "잘못된 API 요청입니다. 누락된 파라미터가 없는지 확인해주세요.",

	// 401 Unauthorized
	ErrInvalidQueryPayload: "JWT 헤더의 페이로드가 올바르지 않습니다. 서명에 사용한 페이로드 값을 확인해주세요.",
	ErrJWTVerification:     "JWT 헤더 검증에 실패했습니다. 토큰이 올바르게 생성, 서명되었는지 확인해주세요.",
	ErrExpiredAccessKey:    "API 키가 만료되었습니다.",
	ErrNonceUsed:           "이미 요청한 nonce값이 다시 사용되었습니다. 매번 새로운 nonce 값을 사용해야 합니다.",
	ErrNoAuthorizationIP:   "허용되지 않은 IP 주소입니다.",
	ErrOutOfScope:          "허용되지 않은 기능입니다.",
}

// GetErrorMessage : 오류 코드에 대응하는 메시지 리턴
// 매핑되지 않은 경우 "Unknown error code" 반환
func GetErrorMessage(code string) string {
	if msg, ok := ErrorMessages[code]; ok {
		return msg
	}
	return "Unknown error code: " + code
}
