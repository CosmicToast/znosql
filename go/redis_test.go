// +build redis

package gozstyle

import (
	"strconv"
	"testing"

	"github.com/go-redis/redis/v7"
)

func BenchmarkPut(b *testing.B) {
	var (
		o, _ = redis.ParseURL("redis://localhost:6379")
		c    = redis.NewClient(o)
	)

	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Set(s, s, 0)
	}
}

func BenchmarkGet(b *testing.B) {
	var (
		o, _ = redis.ParseURL("redis://localhost:6379")
		c    = redis.NewClient(o)
	)

	// populate
	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Set(s, s, 0).Result()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Get(s).Result()
	}
}
