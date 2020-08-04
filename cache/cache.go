package cache

import (
	"fmt"
	"hash/fnv"
	"sync"
)

const (
	bucketInitial = 16
	mapAlloc      = 1024
)

type bucket struct {
	sync.RWMutex
	m map[string][]string
}

type Cache struct {
	buckets []*bucket
	sync.RWMutex
	bucketSize uint32
	entries    uint32
}

func newBucket() *bucket {
	return &bucket{m: make(map[string][]string, mapAlloc)}
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

func (c *Cache) Get(key string) []string {
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

func (c *Cache) rebalance() {
	//stop the world and move the keys
}

// finish doing the stats
func (c *Cache) Stats() {
	c.RLock()
	bucketLen := len(c.buckets) // number of buckets
	mostMapKeys := 0
	mostValuesOneKey := 0
	averageValues := 0
	totalValues := 0
	totalKeys := 0
	averageKeys := 0

	for _, bucket := range c.buckets {
		bucket.RLock()
		curKeys := len(bucket.m)
		for _, val := range bucket.m {
			valLen := len(val)
			if valLen > mostValuesOneKey {
				mostValuesOneKey = valLen
			}
			averageValues += valLen
		}
		bucket.RUnlock()
		if curKeys > mostMapKeys {
			mostMapKeys = curKeys
		}
		averageKeys += curKeys
	}
	c.RUnlock()
	totalKeys = averageKeys
	totalValues += averageValues
	averageValues /= totalKeys
	averageKeys /= bucketLen

	fmt.Println(fmt.Sprintf("Total Keys: %d", totalKeys))
	fmt.Println(fmt.Sprintf("Total Values: %d", totalValues))
	fmt.Println(fmt.Sprintf("Bucket Count: %d", bucketLen))
	fmt.Println(fmt.Sprintf("Average Values per key: %d", averageValues))
	fmt.Println(fmt.Sprintf("Average Keys per map: %d", averageKeys))
	fmt.Println(fmt.Sprintf("Most Keys 1 map: %d", mostMapKeys))
	fmt.Println(fmt.Sprintf("Most Vals 1 key: %d", mostValuesOneKey))
}
