package imcache

import (
	"sync"
	"time"
)

type item struct {
	value interface{}
	ttl   uint64
}

type cache struct {
	data   map[string]*item
	ttl    uint64
	hits   uint64
	misses uint64
	lock   sync.RWMutex
}

func New(ttl uint64) *cache {
	c := cache{
		data:   make(map[string]*item),
		ttl:    ttl,
		hits:   0,
		misses: 0,
	}

	go c.updateTTL()

	return &c
}

func (c *cache) updateTTL() {
	for {
		time.Sleep(time.Second)
		for k, v := range c.data {
			c.lock.Lock()
			v.ttl--
			c.lock.Unlock()
			if v.ttl <= 0 {
				c.lock.Lock()
				delete(c.data, k)
				c.lock.Unlock()
			}
		}
	}
}

func (c *cache) SetTTL(ttl uint64) {
	c.lock.Lock()
	c.ttl = ttl
	c.lock.Unlock()
}

func (c *cache) Set(k string, v interface{}) {
	c.lock.Lock()
	c.data[k] = &item{value: v, ttl: c.ttl}
	c.lock.Unlock()
}

func (c *cache) Get(k string) interface{} {
	c.lock.RLock()
	v, ok := c.data[k]
	c.lock.RUnlock()

	if ok {
		c.hits++
		return v.value
	} else {
		c.misses++
		return nil
	}
}
