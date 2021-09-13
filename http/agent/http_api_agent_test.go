package _agent

import (
	"github.com/no-src/ngo/contract"
	"testing"
)

func TestPost(t *testing.T) {
	url := "https://api.github.com/orgs/ORG/repos/"
	request := _contract.NewApiRequest()
	response, err := Post(url, request)
	if err != nil {
		t.Fail()
		t.Logf("Post 请求失败 err=%s", err.Error())
	} else {
		_ = response
		//t.Logf("Post 成功 response=%v", response)
	}
}
