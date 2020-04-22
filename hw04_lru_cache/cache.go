package hw04_lru_cache //nolint:golint,stylecheck

type Cache interface {
	Set(key string, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key string) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                                 // Очистить кэш
}

type lruCache struct {
	Capacity int              // - capacity
	Queue    *list            // - queue
	Items    map[string]*Item // - items
}

type cacheItem struct {
	Key   string
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		Capacity: capacity,
		Items:    make(map[string]*Item),
		Queue:    &list{},
	}
}

func (c *lruCache) Set(key string, value interface{}) bool {
	// If a key is in the cache, updates it's value and recent-ness
	if element, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(element)
		element.Value.(*cacheItem).Value = value
		return true
	}

	// If cache size is full, remove the least recently used.
	if c.Queue.Len() == c.Capacity {
		if element := c.Queue.Back(); element != nil {
			delete(c.Items, c.Queue.Back().Value.(*cacheItem).Key)
			c.Queue.Remove(c.Queue.Back())
		}
	}
	// If a key not in the cache, adds it and updates it's recent-less
	item := &cacheItem{key, value}
	element := c.Queue.PushFront(item)
	c.Items[item.Key] = element
	return false
}

func (c *lruCache) Get(key string) (interface{}, bool) {
	// If a key is in the cache, returns it's value and updates recent-ness
	if element, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(element)
		return element.Value.(*cacheItem).Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	for key, element := range c.Items {
		delete(c.Items, key)
		c.Queue.Remove(element)
	}
}
