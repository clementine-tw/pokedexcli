package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://example.com",
			val: []byte("test"),
		},
		{
			key: "http://example.com/path",
			val: []byte("path1"),
		},
	}
	interval := 5 * time.Second

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Error("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("vals are different")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {

	interval := 5 * time.Millisecond
	waitTime := 5 * time.Millisecond

	key := "http://example.com"
	val := []byte("testtest")

	cache := NewCache(interval)
	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Error("expected to find key")
		return
	}

	time.Sleep(interval + waitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Error("expected to not find key")
		return
	}
}
