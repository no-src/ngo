package _contract

import (
	"github.com/no-src/ngo/crypto"
	"github.com/no-src/ngo/encoding/json"
)

// ApiRequest 统一请求类
type ApiRequest struct {
	// 是否允许客户端缓存
	AllowClientCache bool
	// 客户端缓存数据的Hash值
	Hash string
	// 客户端业务请求数据
	RequestData interface{}
}

// JSApiRequest 浏览器前端统一请求类
type JSApiRequest struct {
	ApiRequest
	// JSONP客户端回调函数名
	JSONPCallback string
}

// NewApiRequest 创建一个ApiRequest实例
func NewApiRequest() *ApiRequest {
	return &ApiRequest{}
}

// NewJSApiRequest 创建一个JSApiRequest实例
func NewJSApiRequest() *JSApiRequest {
	return &JSApiRequest{}
}

// GetRequestHash 构建请求的Hash键值
func GetRequestHash(url string, requestData interface{}) string {
	requestKey := "API_Request_" + url + "|" + _json.SerializeObject(requestData)
	hash := _crypto.MD5(requestKey)
	return hash
}
