package cache

import (
	"testing"
	"time"
)

func BenchmarkCacheOperations(b *testing.B) {
	c := NewCustomCache()
	key := []byte("key")
	value := []byte("value")
	ttl := 10 * time.Second

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set(key, value, ttl)
		c.Get(key)
	}
}
