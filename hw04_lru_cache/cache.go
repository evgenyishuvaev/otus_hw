package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	item, ok := cache.items[key]

	if cache.isFull() {
		cache.deleteLRUItem()
	}

	if ok {
		item.Value = value
		cache.queue.MoveToFront(item)
	} else {
		item := cache.queue.PushFront(value)
		cache.items[key] = item
	}
	return ok
}

func (cache *lruCache) deleteLRUItem() {
	lastItem := cache.queue.Back()
	for key, val := range cache.items {
		if val.Value == lastItem.Value {
			delete(cache.items, key)
			break
		}
	}
	cache.queue.Remove(lastItem)
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := cache.items[key]
	if !ok {
		return nil, false
	}
	cache.queue.MoveToFront(item)
	return item.Value, ok
}

func (cache *lruCache) isFull() bool {
	return cache.queue.Len() == cache.capacity
}

func (cache *lruCache) Clear() {
	for key, item := range cache.items {
		cache.queue.Remove(item)
		delete(cache.items, key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
