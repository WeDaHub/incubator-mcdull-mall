package dao

import (
	"log"
	"sync"
	"time"
)

var cache = sync.Map{}
var cacheTTL = sync.Map{}

// 清理过期的缓存
func walk(key, value interface{}) bool {
	if time.Now().Unix()-value.(int64) > 0 {
		cache.Delete(key)
		cacheTTL.Delete(key)
		log.Printf("Clean up expired cache key = %s, ttl = %v\n", key, value)
	}
	return true
}

func init() {
	ticker := time.NewTicker(10 * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			log.Printf("ticker trigger %s", time.Now())
			cacheTTL.Range(walk)
		}
	}(ticker)
}

// GetLocalCache 获取本地缓存
func GetLocalCache(cacheKey string) (interface{}, bool) {
	if ttl, ok := cacheTTL.Load(cacheKey); ok {
		if time.Now().Unix()-ttl.(int64) > 0 {
			return nil, false
		}
	}
	if cacheValue, ok := cache.Load(cacheKey); ok {
		return cacheValue, true
	}
	return nil, false
}

// SetLoadCache 设置本次缓存 + TTL值（单位：s）
func SetLoadCache(cacheKey string, cacheValue interface{}, ttl uint32) {
	cache.Store(cacheKey, cacheValue)
	cacheTTL.Store(cacheKey, time.Now().Unix()+int64(ttl))
}

// CleanLocalCache 清理缓存
func CleanLocalCache(cacheKey string) {
	cache.Delete(cacheKey)
	cacheTTL.Delete(cacheKey)
}
