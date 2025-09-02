package utilities

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(length int) (string, error) {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charsetLen := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		randomString[i] = charset[index.Int64()]
	}
	return string(randomString), nil
}
