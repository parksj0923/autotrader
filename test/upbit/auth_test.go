package upbit

import (
	"autotrader/main/config"
	"autotrader/main/domain/service/exchange"
	"fmt"
	"testing"
)

func TestKeys(t *testing.T) {
	accessKey := config.DefaultEnv.AccessKey
	secretKey := config.DefaultEnv.SecretKey
	fmt.Println(accessKey, secretKey)
}

func TestGetJwtTokenWithNilParams(t *testing.T) {
	token, err := exchange.GenerateJWT(nil)
	if err != nil {
		return
	}
	fmt.Println(token)
}
