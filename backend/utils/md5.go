package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 用于Token编码
func Md5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
