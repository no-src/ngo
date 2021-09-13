package _serialize

// Serializer 序列化器接口
type Serializer interface {
	// Marshal 将对象序列化为byte数组
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal 将byte数组反序列化为对象
	Unmarshal(data []byte, v interface{}) error
}
