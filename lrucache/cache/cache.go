package cache

// HashMap is a map of strings to Nodes.
type HashMap map[string]*Node

// Cache is a simple LRU cache.
type Cache struct {
	Queue   *Queue
	HashMap *HashMap
}

// NewCache creates a new Cache.
func NewCache(size int) *Cache {
	return &Cache{Queue: NewQueue(size), HashMap: &HashMap{}}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) string {
	if node, ok := (*c.HashMap)[key]; ok {
		c.Queue.MoveToFront(node)
		return node.value
	}

	return ""
}

// Set sets a value in the cache
func (c *Cache) Set(key, value string) {
	if node, ok := (*c.HashMap)[key]; ok {
		node.value = value
		c.Queue.MoveToFront(node)
		return
	}

	c.Queue.Add(key, value)
	(*c.HashMap)[key] = c.Queue.head.next
}

// Remove removes a value from the cache
func (c *Cache) Remove(key string) {
	if node, ok := (*c.HashMap)[key]; ok {
		c.Queue.Remove(node)
		delete(*c.HashMap, key)
	}
}
