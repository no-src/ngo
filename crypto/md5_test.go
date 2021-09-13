package _crypto

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	hash := MD5("111111")
	expect := "96e79218965eb72c92a549dd5a330112"
	if hash != expect {
		fmt.Printf("md5 calc error ,expect=%s but get %s \r\n", expect, hash)
		t.Fail()
	}
}
