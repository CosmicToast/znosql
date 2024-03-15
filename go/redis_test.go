//+build redis

package gozstyle

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/redis/go-redis/v9"
)

func TestKV(t *testing.T) {
	assert := assert.New(t)
	const (
		k  = "key"
		v1 = "val1"
		v2 = "val2"
	)
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	assert.NotNil(c)
	assert.Nil(c.Ping(ctx).Err())

	assert.Nil(c.Set(ctx, k, v1, 0).Err())
	o1 := c.Get(ctx, k)
	assert.Nil(o1.Err())
	assert.Equal(v1, o1.Val())

	assert.Nil(c.Set(ctx, k, v2, 0).Err())
	o2 := c.Get(ctx, k)
	assert.Nil(o2.Err())
	assert.Equal(v2, o2.Val())

	c.Close()
}

func BenchmarkPut(b *testing.B) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Set(ctx, s, s, 0)
	}
}

func BenchmarkGet(b *testing.B) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// populate
	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Set(ctx, s, s, 0).Result()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Get(ctx, s).Result()
	}
}
