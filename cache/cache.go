package cache

import (
	"hash/fnv"
	"sync"
)

const (
	bucketInitial = 16
)

type bucket struct {
	sync.RWMutex
	m map[string][]string
}

type Cache struct {
	buckets []*bucket
	sync.RWMutex
	bucketSize uint32
}

func newBucket() *bucket {
	return &bucket{m: make(map[string][]string, bucketInitial)}
}

func NewCache() *Cache {
	buckets := make([]*bucket, bucketInitial)
	for ind := range buckets {
		buckets[ind] = newBucket()
	}
	return &Cache{buckets: buckets, bucketSize: bucketInitial}
}
func (c *Cache) Insert(key string, value string) {
	bucket := c.bucketForKey(key)
	bucket.Lock()
	defer bucket.Unlock() // TODO save some cycles no defer
	val, ok := bucket.m[key]
	if !ok {
		val = []string{value}
		bucket.m[key] = val
		return
	}
	val = append(val, value)
	bucket.m[key] = val // required as position may move when resizing slice?
	// don't need to reenter in the map since the reference should not change
}

func (c *Cache) Get(key string) interface{} {
	bucket := c.bucketForKey(key)
	return bucket.m[key]
}

// idempotent
func (c *Cache) bucketForKey(key string) *bucket {
	hashf := fnv.New32() // TODO benchmark pooling these, make sure you can actually just spawn functions and they're idempotent (:
	hashf.Write([]byte(key))
	bucket := hashf.Sum32() % c.bucketSize
	c.RLock()
	m := c.buckets[bucket]
	c.RUnlock()
	return m
}
