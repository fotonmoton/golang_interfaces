package cache

type Evictable interface {
	evict(c *Cache)
}

type Cache struct {
	storage     map[string]string
	eviction    Evictable
	size        int
	maxCapacity int
}

func NewCache(e Evictable) *Cache {
	return &Cache{
		storage:     make(map[string]string),
		eviction:    e,
		size:        0,
		maxCapacity: 1,
	}
}

func (c *Cache) Add(key, value string) {
	if c.size == c.maxCapacity {
		c.eviction.evict(c)
		c.size--
	}
	c.size++
	c.storage[key] = value
}

func (c *Cache) Get(key string) string {
	return c.storage[key]
}
