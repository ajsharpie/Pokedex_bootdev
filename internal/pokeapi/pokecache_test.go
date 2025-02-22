package pokeapi

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(2)

	key := "testKey"
	value := []byte("testValue")

	cache.CacheAdd(key, value)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected to find key %s", key)
	}
	if string(got) != string(value) {
		t.Fatalf("expected value %s, got %s", value, got)
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(1)

	key := "testKey"
	value := []byte("testValue")

	cache.CacheAdd(key, value)

	time.Sleep(2 * time.Second)

	_, ok := cache.Get(key)
	if ok {
		t.Fatalf("expected key %s to be expired", key)
	}
}

func TestCacheOverwrite(t *testing.T) {
	cache := NewCache(2)

	key := "testKey"
	value1 := []byte("testValue1")
	value2 := []byte("testValue2")

	cache.CacheAdd(key, value1)
	cache.CacheAdd(key, value2)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected to find key %s", key)
	}
	if string(got) != string(value2) {
		t.Fatalf("expected value %s, got %s", value2, got)
	}
}
