package tbase64

import (
	"encoding/base64"
)

type TBase64 struct {
}

// Encode base64加密
func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Decode 解密
func Decode(s string) (string, error) {
	ds, err := base64.StdEncoding.DecodeString(s)
	return string(ds), err
}
