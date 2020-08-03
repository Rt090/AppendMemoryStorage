package cache

type Cache struct{}

func (c *Cache) Insert(key string, value interface{}) {}
func (c *Cache) Get(key string) interface{} {
	return nil
}
