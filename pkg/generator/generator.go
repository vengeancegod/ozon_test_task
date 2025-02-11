package generator

import (
	"crypto/rand"
	"math/big"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = 10
)

func GenerateShortURL() string {
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		b[i] = alphabet[num.Int64()]
	}
	return string(b)
}
