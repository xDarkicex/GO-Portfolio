package helpers

import (
	"sync"
	"time"
)

// CacheObject represents a cached object with expiration time.
type CacheObject struct {
	Expire time.Time
	Object interface{}
}

// Cache defines the contract for a cache.
type Cache interface {
	Get(key string, fallback func() *CacheObject) *CacheObject
	Delete(key string)
	SetCleanupInterval(interval time.Duration)
}

// InMemoryCache is a simple in-memory cache with expiration.
type InMemoryCache struct {
	cache             map[string]*CacheObject
	mutex             sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

// NewInMemoryCache initializes a new in-memory cache with the given default expiration duration and cleanup interval.
func NewInMemoryCache(defaultExpiration, cleanupInterval time.Duration) *InMemoryCache {
	cache := &InMemoryCache{
		cache:             make(map[string]*CacheObject),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}
	go cache.periodicCleanup()
	return cache
}

// Get retrieves a cached object by key. If not found, it invokes the fallback function to compute the object.
func (c *InMemoryCache) Get(key string, fallback func() *CacheObject) *CacheObject {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if obj, ok := c.cache[key]; ok {
		return obj
	}

	newObject := fallback()
	c.cache[key] = newObject
	return newObject
}

// Delete removes a cached object by key.
func (c *InMemoryCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cache, key)
}

// SetCleanupInterval sets a new cleanup interval for the cache.
func (c *InMemoryCache) SetCleanupInterval(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cleanupInterval = interval
}

// periodicCleanup periodically purges expired items from the cache.
func (c *InMemoryCache) periodicCleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, cachedObject := range c.cache {
			if time.Now().After(cachedObject.Expire) {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}


/* 
// original code
package helpers

import (
	"sync"
	"time"
)

// CacheObject shut up
type CacheObject struct {
	Expire time.Time
	Object interface{}
}

// Cache Actual cache in memory
var cache = make(map[string]*CacheObject)
var mutex = &sync.Mutex{}

const defaultExpiration = time.Hour * time.Duration(1)

// Get cached Object
// key: the key our cached object.
// fallback: code to execute to compute our cached object. Or whatever.
//
// Example:
// object := Cache.Get("layout", func() {
//   getOurLayout := parseFileOrWhatever();
//   return NewCacheObject(getOurLayoutShit)
// })
//

func Get(key string, fallback func() *CacheObject) *CacheObject {
	mutex.Lock()
	if cache[key] == nil {
		cache[key] = fallback()
	}
	mutex.Unlock()
	return cache[key]

}

// DeleteCache object for updating pages
func DeleteCache(key string) {
	delete(cache, key)
}

// NewCacheObject Creates Pointer to new object
func NewCacheObject(object interface{}) *CacheObject {
	return &CacheObject{
		Expire: time.Now().Add(defaultExpiration),
		Object: object,
	}
}

func init() {
	var timer *time.Timer
	timer = time.AfterFunc(defaultExpiration, func() {
		for key, cachedObject := range cache {
			if time.Now().After(cachedObject.Expire) {
				delete(cache, key)
			}
		}
		timer.Reset(defaultExpiration)
	})

}
*/
