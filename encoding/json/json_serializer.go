package _json

import (
	"encoding/json"
)

// JsonSerializer JSON序列化器
type JsonSerializer struct {
}

// Marshal 将对象序列化为JSON字符串
func (current JsonSerializer) Marshal(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal 将JSON串反序列化为对象
func (current JsonSerializer) Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}
