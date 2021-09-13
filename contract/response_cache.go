package _contract

import "github.com/no-src/ngo/cache"

// ApiResponseCache 响应数据缓存Model
type ApiResponseCache struct {
	_cache.CacheData
	//数据
	Data ApiResponse
}
