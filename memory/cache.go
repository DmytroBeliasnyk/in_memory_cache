package memory

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
