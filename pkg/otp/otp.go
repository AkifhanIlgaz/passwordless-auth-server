package otp

import (
	"math/rand"
	"time"
)

func CreateCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = digits[r.Intn(len(digits))]
	}

	return string(code)
}
