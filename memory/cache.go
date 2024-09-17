package memory

import "errors"

var single *cache

type cache struct {
	memoryCache map[string]interface{}
}

func GetCache() *cache {
	if single == nil {
		single = &cache{
			memoryCache: make(map[string]interface{}),
		}
	}

	return single
}

func (c cache) Set(key string, value interface{}) error {
	_, exist := single.memoryCache[key]
	if exist {
		return errors.New("the key " + key + " already exists")
	}

	single.memoryCache[key] = value
	return nil
}

func (c cache) Get(key string) (interface{}, error) {
	res, exist := single.memoryCache[key]
	if !exist {
		return 0, errors.New("no mapping for the key " + key)
	}

	return res, nil
}

func (c cache) Delete(key string) {
	delete(single.memoryCache, key)
}
