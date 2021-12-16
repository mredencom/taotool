package tmd5

import (
	"crypto/md5"
	"encoding/hex"
)

type TMd5 struct {
}

// Md5 加密
func (TMd5) Md5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
