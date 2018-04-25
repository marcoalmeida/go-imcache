package imcache

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cache := New(3)

	if len(cache.data) != 0 {
		t.Error("Expected no data, got", cache.data)
	}

	if cache.ttl != 3 {
		t.Error("Wrong TTL")
	}

	if cache.misses != 0 {
		t.Error("Expected no cache misses, got", cache.misses)
	}

	if cache.hits != 0 {
		t.Error("Expected no cache hits, got", cache.hits)
	}
}

func TestCache_SetTTL(t *testing.T) {
	cache := New(3)

	if cache.ttl != 3 {
		t.Error("Wrong TTL")
	}

	cache.SetTTL(6)

	if cache.ttl != 6 {
		t.Error("Wrong TTL")
	}
}

func TestCache_Set(t *testing.T) {
	cache := New(3)

	k := "key"
	v0 := "value"
	v1 := "update"

	cache.Set(k, v0)
	value := cache.data[k]
	if value.value != v0 {
		t.Error("Expected ", v0, "got", value.value)
	}

	// update
	cache.Set(k, v1)
	value = cache.data[k]
	if value.value != v1 {
		t.Error("Expected ", v1, "got", value.value)
	}
}

func TestCache_Get(t *testing.T) {
	cache := New(3)

	k := "key"
	v := "value"

	cache.Set(k, v)

	value := cache.Get(k)
	if value != v {
		t.Error("Expected", v, "got", value)
	}

	// expire the TTL
	time.Sleep(4 * time.Second)

	value = cache.Get(k)
	if value != nil {
		t.Error("Expected", nil, "got", value)
	}
}
