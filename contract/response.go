package _contract

import (
	"github.com/no-src/ngo/crypto"
	"github.com/no-src/ngo/encoding/json"
)

// ApiResponse 统一响应结果
type ApiResponse struct {
	// 服务是否出错
	HasError bool
	// 响应代码，0表示正常，小于0为服务出错代码，大于0为业务代码
	Code int
	// 响应代码对应的提示消息
	Message string
	// 服务返回的数据
	Data interface{}
	// Data的唯一Hash值
	Hash string
	// 客户端是否已经缓存,并且与服务器的数据一致
	ClientCached bool
}

// NewApiResponse 创建一个ApiResponse实例
func NewApiResponse() *ApiResponse {
	return &ApiResponse{}
}

// GetResponseHash 获取响应数据的哈希值
func GetResponseHash(responseMessage interface{}) string {
	hash := _crypto.MD5(_json.SerializeObject(responseMessage))
	return hash
}

// SupportClientCache 处理ApiResponse，如果客户端启用缓存，并且客户端的数据的Hash与服务端一致，则不响应给客户端具体的Data
func SupportClientCache(request *ApiRequest, response *ApiResponse) {
	if request == nil || response == nil || request.AllowClientCache == false {
		return
	}
	hash := GetResponseHash(response.Data)
	response.Hash = hash
	if request.Hash == hash && hash != "" {
		response.ClientCached = true
		response.Data = nil
	}
}
