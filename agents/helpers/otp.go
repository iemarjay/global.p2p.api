package helpers

import (
	"crypto/rand"
	"global.p2p.api/app"
	"global.p2p.api/app/cache"
	"math/big"
	"time"
)

type OtpGenerator struct {
	cache *cache.Cache
}

func (g OtpGenerator) token(length int) (string, error) {
	seed := "0123456789"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}

func (g OtpGenerator) TokenFor(key string) (string, error) {
	token, err := g.token(6)

	if err != nil {
		return "", err
	}

	err = g.cache.Set(key, token, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (g OtpGenerator) Validate(key string, value string) bool {
	token, err := g.cache.Get(key)
	if err != nil {
		return false
	}

	return token == value
}

func Otp(env *app.Env) *OtpGenerator {
	return &OtpGenerator{cache: cache.New(env)}
}
