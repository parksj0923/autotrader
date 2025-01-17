package config

import "os"

type Exchange string

const (
	UPBIT Exchange = "upbit"
)

type CommonEnv struct {
	Exchange  Exchange
	AccessKey string
	SecretKey string
}

var DefaultEnv = SetUpEnv()

func SetUpEnv() CommonEnv {
	if Exchange(os.Getenv("EXCHANGE")) == UPBIT {
		return CommonEnv{
			Exchange:  UPBIT,
			AccessKey: os.Getenv("UPBIT_ACCESS_KEY"),
			SecretKey: os.Getenv("UPBIT_SECRET_KEY"),
		}
	}
	return CommonEnv{}
}
