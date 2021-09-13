package _cache

// Cache 缓存抽象接口
type Cache interface {
	// Get 获取缓存数据
	Get(key string) (interface{}, error)
	// Set 写入缓存数据，过期时间单位为秒
	Set(key string, data interface{}, cacheSeconds int) error
	// Delete 删除指定的缓存key
	Delete(key string) error
}
