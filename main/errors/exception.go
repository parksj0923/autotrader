package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type errorBase struct {
	Reason string `json:"reason"`
	Status int    `json:"status"`
	Msg    string `json:"data"`
}

func (base errorBase) Error() string {
	return base.Msg
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (base errorBase) NewErrorResponse() ErrorResponse {
	return ErrorResponse{Code: base.Reason, Message: base.Msg}
}

func UnMarshalErrorResponse(errorString string) (ErrorResponse, error) {
	var base ErrorResponse
	err := json.Unmarshal([]byte(errorString), &base)
	return base, err
}

func ConvertToErrorBase(err error) (errorBase, error) {
	var convertedError *errorBase
	converted := errors.As(err, &convertedError)
	if converted {
		return *convertedError, nil
	}
	return errorBase{}, err
}

func NewRequestParserError(requestType string) error {
	return &errorBase{
		Reason: "RequestParseError",
		Status: fiber.StatusBadRequest,
		Msg:    fmt.Sprintf("request json parse error struct: %s", requestType),
	}
}

func NewCustomError(reason string, status int, msg string) error {
	return &errorBase{
		Reason: reason,
		Status: status,
		Msg:    msg,
	}
}

func NewLoginFailed(email string) error {
	return &errorBase{
		Reason: "LoginFailed",
		Status: fiber.StatusBadRequest,
		Msg:    fmt.Sprintf("The requested login info is not correct : %s", email),
	}
}

func NewTokenExpired() error {
	return &errorBase{
		Reason: "TokenExpired",
		Status: fiber.StatusUnauthorized,
		Msg:    fmt.Sprintf("The token in the request header is expired."),
	}
}

func NewTokenInvalid(jwt string) error {
	return &errorBase{
		Reason: "TokenInvalid",
		Status: fiber.StatusUnauthorized,
		Msg:    fmt.Sprintf("The token of the header doesn't have valid information : %s", jwt),
	}
}

func NewTokenReplaced() error {
	return &errorBase{
		Reason: "TokenReplaced",
		Status: fiber.StatusUnauthorized,
		Msg:    fmt.Sprintf("This token has been replaced with a newer token."),
	}
}

func NewUnauthorized() error {
	return &errorBase{
		Reason: "Unauthorized",
		Status: fiber.StatusUnauthorized,
		Msg:    fmt.Sprintf("The request doesn't have a authorization token"),
	}
}

func NewInternalServerError() error {
	return &errorBase{
		Reason: "Internal Server Error",
		Status: fiber.StatusInternalServerError,
		Msg:    "Internal Server Error",
	}
}

func NewTooManyRequestForRateLimit(key string) error {
	return &errorBase{
		Reason: "Too many Requests",
		Status: fiber.StatusTooManyRequests,
		Msg:    key,
	}
}
