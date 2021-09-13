package _cache

// EmptyCache 空的缓存实现，不做任何操作
type EmptyCache struct {
}

func (cache *EmptyCache) Get(key string) (interface{}, error) {
	return nil, nil
}

func (cache *EmptyCache) Set(key string, data interface{}, cacheSeconds int) error {
	return nil
}

func (cache *EmptyCache) Delete(key string) error {
	return nil
}
