package cache

import (
	"sync"
	"time"
)

type node struct {
	key       []byte
	value     []byte
	expiry    time.Time
	next      *node
	prev      *node
}

type CustomCache struct {
	head, tail *node
	lock       sync.Mutex
}

func NewCustomCache() *CustomCache {
	c := &CustomCache{}
	go c.cleanupExpiredKeys()
	return c
}

func (c *CustomCache) Set(key, value []byte, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// Create new node
	newNode := &node{
		key:    append([]byte(nil), key...),
		value:  append([]byte(nil), value...),
		expiry: time.Now().Add(ttl),
	}

	// Append to the tail of the list
	if c.tail == nil {
		c.head = newNode
		c.tail = newNode
	} else {
		c.tail.next = newNode
		newNode.prev = c.tail
		c.tail = newNode
	}
}

func (c *CustomCache) Get(key []byte) (value []byte, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	current := c.head
	for current != nil {
		if time.Now().After(current.expiry) {
			current = current.next
			continue
		}
		if string(current.key) == string(key) {
			ttl = time.Until(current.expiry)
			return append([]byte(nil), current.value...), ttl
		}
		current = current.next
	}

	return nil, 0
}

func (c *CustomCache) cleanupExpiredKeys() {
	for {
		time.Sleep(1 * time.Minute) // Cleanup interval
		c.lock.Lock()
		for c.head != nil && time.Now().After(c.head.expiry) {
			c.head = c.head.next
			if c.head != nil {
				c.head.prev = nil
			}
		}
		if c.head == nil {
			c.tail = nil
		}
		c.lock.Unlock()
	}
}
