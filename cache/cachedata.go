package _cache

import "time"

type CacheData struct {

	//创建时间
	CreateTime time.Time

	//数据缓存过期时间
	CacheExpireTime time.Time

	// 数据实际存储过期时间
	StoreExpireTime time.Time

	//数据
	Data interface{}
}
