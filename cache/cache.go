package cache

import (
	"sync"
	"time"
)

// node defines the structure of each list element.
type node struct {
	key    []byte
	value  []byte
	expiry time.Time
	next   *node
	prev   *node
}

// CustomCache maintains the head and tail pointers and includes a mutex for synchronization.
type CustomCache struct {
	head, tail *node
	lock       sync.Mutex
}

// NewCustomCache initializes and returns a new cache instance.
func NewCustomCache() *CustomCache {
	c := &CustomCache{}
	go c.cleanupExpiredKeys() // Start background task to clean up expired keys.
	return c
}

// Set appends a new node with key, value, and expiry to the linked list.
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

// Get searches for the key in the list and returns the value and remaining TTL if not expired.
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

// cleanupExpiredKeys removes nodes from the list that have expired, running in a separate goroutine.
func (c *CustomCache) cleanupExpiredKeys() {
	for {
		time.Sleep(1 * time.Minute) // Interval to check and remove expired nodes.
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
