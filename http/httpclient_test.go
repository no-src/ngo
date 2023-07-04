package _http

import (
	"testing"

	"github.com/no-src/log"
	"github.com/no-src/log/level"
	"github.com/no-src/ngo/strings"
)

func TestGet(t *testing.T) {
	defer log.Close()
	url := "https://github.com/"
	httpClient := NewHttpClient()
	responseData, err := httpClient.Get(url, nil, _strings.Empty, nil, nil)
	if err != nil {
		log.Error(err, "Get Request Failed:%s", url)
		t.Fail()
	} else {
		content := string(responseData)
		_ = content
		//log.Debug("Get Request:%s,Result:%s", url, content)
	}
}

func TestPostJson(t *testing.T) {
	defer log.Close()
	url := "https://github.com/"
	httpClient := NewHttpClient()
	responseData, err := httpClient.PostJson(url, nil)
	if err != nil {
		log.Error(err, "Post Request Failed：%s", url)
		t.Fail()
	} else {
		content := string(responseData)
		_ = content
		//log.Debug("Post Request:%s,Result:%s", url, content)
	}
}

func init() {
	log.InitDefaultLogger(log.NewConsoleLogger(level.DebugLevel))
}
