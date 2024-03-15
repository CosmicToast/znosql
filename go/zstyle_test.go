//+build zstyle

package gozstyle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKV(t *testing.T) {
	assert := assert.New(t)
	const (
		k  = "key"
		v1 = "val1"
		v2 = "val2"
	)

	c := NewZstyle("localhost:5050")
	assert.NotNil(c)
	assert.True(c.Ping())

	c.Put(k, v1)
	v, e := c.Get(k)
	assert.Nil(e)
	assert.Equal(v1, v)

	c.Put(k, v2)
	v, e = c.Get(k)
	assert.Nil(e)
	assert.Equal(v2, v)

	c.Exit() // NEVER FORGET TO EXIT
}

func BenchmarkPut(b *testing.B) {
	c := NewZstyle("localhost:5050")

	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Put(s, s)
	}

	c.Exit()
}

func BenchmarkGet(b *testing.B) {
	c := NewZstyle("localhost:5050")

	// populate
	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Put(s, s)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := strconv.Itoa(n)
		c.Get(s)
	}

	c.Exit()
}
