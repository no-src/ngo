package _crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 计算字符串的MD5值
func MD5(str string) (hash string) {
	md5Provider := md5.New()
	md5Provider.Write([]byte(str))
	sum := md5Provider.Sum(nil)
	hash = hex.EncodeToString(sum)
	return hash
}
