package cache

import (
	"testing"
	"time"
)

// BenchmarkCacheOperations tests the Set and Get operations of the cache.
func BenchmarkCacheOperations(b *testing.B) {
	c := NewCustomCache()
	key := []byte("key")
	value := []byte("value")
	ttl := 10 * time.Second
	
	// Reports memory allocations which are part of the benchmarking.
	b.ReportAllocs() 
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set(key, value, ttl)
		c.Get(key)
	}
}
