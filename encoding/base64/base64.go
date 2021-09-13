package _base64

import "encoding/base64"

//Base64Encode Base64编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//Base64Decode Base64解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

//Base64EncodeStr Base64编码
func Base64EncodeStr(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

//Base64DecodeStr Base64解码
func Base64DecodeStr(src string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
