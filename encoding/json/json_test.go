package _json

import (
	"fmt"
	"testing"
)

type TestUserInfo struct {
	UserName string
	Age      int
}

var userInfo = TestUserInfo{
	UserName: "test_user",
	Age:      21,
}

var userInfoJson = `{"UserName":"test_user","Age":21}`

func TestMarshal(t *testing.T) {
	jsonData, err := Marshal(userInfo)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(string(jsonData))
	}
}

func TestUnmarshal(t *testing.T) {
	p := new(TestUserInfo)
	err := Unmarshal([]byte(userInfoJson), &p)
	if err != nil || p.UserName != "test_user" {
		t.Fail()
	}
}

func TestSerializeObject(t *testing.T) {
	jsonData := SerializeObject(userInfo)
	if len(jsonData) == 0 {
		t.Fail()
	} else {
		fmt.Println(string(jsonData))
	}
}

func TestDeserializeObject(t *testing.T) {
	p := new(TestUserInfo)
	err := DeserializeObject(userInfoJson, &p)
	if err != nil || p.UserName != "test_user" {
		t.Fail()
	}
}
