package _base64

import "testing"

func TestBase64Encode(t *testing.T) {
	testBase64Encode(t, "hello", "aGVsbG8=")
	testBase64Encode(t, "helloworld", "aGVsbG93b3JsZA==")
	testBase64Encode(t, "hello世界", "aGVsbG/kuJbnlYw=")
	testBase64Encode(t, "再见", "5YaN6KeB")
}

func testBase64Encode(t *testing.T, src, expect string) {
	bytes := []byte(src)
	result := Base64Encode(bytes)
	if result != expect {
		t.Fail()
		t.Logf("TestBase64Encode Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestBase64Decode(t *testing.T) {
	testBase64Decode(t, "aGVsbG8=", "hello")
	testBase64Decode(t, "aGVsbG93b3JsZA==", "helloworld")
	testBase64Decode(t, "aGVsbG/kuJbnlYw=", "hello世界")
	testBase64Decode(t, "5YaN6KeB", "再见")
}

func testBase64Decode(t *testing.T, src, expect string) {
	resultBytes, err := Base64Decode(src)
	if err != nil {
		t.Logf("TestBase64Encode Failed src=[%s] err=%s \r\n", src, err.Error())
	}
	result := string(resultBytes)
	if result != expect {
		t.Fail()
		t.Logf("TestBase64Encode Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestBase64EncodeStr(t *testing.T) {
	testBase64EncodeStr(t, "hello", "aGVsbG8=")
	testBase64EncodeStr(t, "helloworld", "aGVsbG93b3JsZA==")
	testBase64EncodeStr(t, "hello世界", "aGVsbG/kuJbnlYw=")
	testBase64EncodeStr(t, "再见", "5YaN6KeB")
}

func testBase64EncodeStr(t *testing.T, src, expect string) {
	result := Base64EncodeStr(src)
	if result != expect {
		t.Fail()
		t.Logf("TestBase64Encode Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestBase64DecodeStr(t *testing.T) {
	testBase64DecodeStr(t, "aGVsbG8=", "hello")
	testBase64DecodeStr(t, "aGVsbG93b3JsZA==", "helloworld")
	testBase64DecodeStr(t, "aGVsbG/kuJbnlYw=", "hello世界")
	testBase64DecodeStr(t, "5YaN6KeB", "再见")
}

func testBase64DecodeStr(t *testing.T, src, expect string) {
	result, err := Base64DecodeStr(src)
	if err != nil {
		t.Logf("TestBase64Encode Failed src=[%s] err=%s \r\n", src, err.Error())
	}
	if result != expect {
		t.Fail()
		t.Logf("TestBase64Encode Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}
