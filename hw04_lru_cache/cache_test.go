package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)
		_, ok := c.Get("aaa")
		require.False(t, ok)
		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)
		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)
		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})
	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(2)
		c.Set("1", "aaa")
		c.Set("2", "bbb")
		c.Set("3", "ccc")

		val, ok := c.Get("1")
		require.False(t, ok)
		val, ok = c.Get("3")
		require.True(t, ok)
		require.Equal(t, "ccc", val)
		wasInCache := c.Set("1", "ddd")
		require.False(t, wasInCache)

		val, ok = c.Get("3")
		require.True(t, ok)
		require.Equal(t, "ccc", val)

		val, ok = c.Get("2")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("clear cache", func(t *testing.T) {
		c := NewCache(2)
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Clear()

		val, ok := c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("Get function updates recent-ness", func(t *testing.T) {
		c := NewCache(2)
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Get("aaa")
		c.Set("ccc", 300)

		val, ok := c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("Set function updates recent-ness", func(t *testing.T) {
		c := NewCache(2)
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("aaa", 200)
		c.Set("ccc", 300)

		val, ok := c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

/* func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
} */
