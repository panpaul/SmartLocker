package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func EncodeSha256(key string) string {
	s := sha256.New()
	s.Write([]byte(key))
	return hex.EncodeToString(s.Sum(nil))
}

func RandString(len int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
