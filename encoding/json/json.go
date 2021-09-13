package _json

import (
	"github.com/no-src/ngo/encoding/serialize"
	"github.com/no-src/ngo/strings"
)

var (
	// 内部JSON序列化器
	jsonSerializer _serialize.Serializer
)

// Marshal 将对象序列化为JSON字符串
func Marshal(v interface{}) ([]byte, error) {
	return jsonSerializer.Marshal(v)
}

// Unmarshal 将JSON串反序列化为对象
func Unmarshal(data []byte, v interface{}) error {
	return jsonSerializer.Unmarshal(data, v)
}

// SerializeObject 将对象序列化为JSON字符串,如果发生异常则返回空字符串
func SerializeObject(v interface{}) string {
	data, err := jsonSerializer.Marshal(v)
	if err != nil {
		return _strings.Empty
	}
	return string(data)
}

// DeserializeObject 将JSON串反序列化为对象
func DeserializeObject(jsonStr string, v interface{}) error {
	err := jsonSerializer.Unmarshal([]byte(jsonStr), &v)
	return err
}

// SetSerializer 重置JSON序列化器
func SetSerializer(serializer _serialize.Serializer) {
	jsonSerializer = serializer
}

// initDefaultJsonSerializer 初始化默认JSON序列化器
func initDefaultJsonSerializer() {
	jsonSerializer = JsonSerializer{}
}

func init() {
	initDefaultJsonSerializer()
}
