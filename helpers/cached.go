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
