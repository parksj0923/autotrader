package utils

import (
	"autotrader/main/config"
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT : Upbit용 JWT 생성
// - accessKey: 업비트 Access Key
// - secretKey: 업비트 Secret Key
// - params: 쿼리 파라미터 (없을 경우 nil 또는 빈 map 전달)
func GenerateJWT(params map[string]string) (string, error) {
	accessKey := config.DefaultEnv.AccessKey
	secretKey := config.DefaultEnv.SecretKey
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + strconv.Itoa(rand.Intn(100000))
	claims := jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      nonce,
	}

	if len(params) > 0 {
		queryString := ""
		for k, v := range params {
			queryString += k + "=" + v + "&"
		}
		queryString = queryString[:len(queryString)-1] // 마지막 `&` 제거

		hash := sha512.New()
		hash.Write([]byte(queryString))
		queryHash := hex.EncodeToString(hash.Sum(nil))

		claims["query_hash"] = queryHash
		claims["query_hash_alg"] = "SHA512"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
