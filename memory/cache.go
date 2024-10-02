package memory

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var single *Cache

type Cache struct {
	memoryCache map[string]interface{}
	mu          sync.RWMutex
}

type item struct {
	key string
	ttl time.Duration
}

func GetCache() *Cache {
	if single == nil {
		single = &Cache{
			memoryCache: make(map[string]interface{}),
			mu:          sync.RWMutex{},
		}
	}

	return single
}

func (c *Cache) String() (res string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	res = "IN MEMORY CACHE:\n"
	for k, v := range c.memoryCache {
		res += fmt.Sprintf("#%s - %v\n", k, v)
	}

	return
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	defer c.mu.Unlock()
	c.mu.Lock()

	_, exist := single.memoryCache[key]

	switch {
	case exist:
		return fmt.Errorf("the key: %s - already exists", key)
	case ttl <= 0:
		return fmt.Errorf("invalid time-to-live value: %s", ttl.String())
	}

	single.memoryCache[key] = value

	items := make(chan item)
	go c.ttl(items)

	items <- item{key, ttl}
	close(items)

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	res, exist := single.memoryCache[key]
	if !exist {
		return 0, errors.New("no mapping for the key " + key)
	}

	return res, nil
}

func (c *Cache) Delete(key string) {
	defer c.mu.Unlock()
	c.mu.Lock()

	delete(single.memoryCache, key)
}

func (c *Cache) ttl(items <-chan item) {
	timer := func(key string, t time.Duration, wg *sync.WaitGroup) {
		<-time.After(t)
		c.Delete(key)
		wg.Done()
	}

	wg := &sync.WaitGroup{}
	for i := range items {
		wg.Add(1)
		go timer(i.key, i.ttl, wg)
	}

	wg.Wait()
}
