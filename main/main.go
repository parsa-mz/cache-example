package main

import (
    "fmt"
    "time"
    "strconv"
    "cache-example/cache" 
)

func main() {
    // Initialize the cache from the cache package.
    c := cache.NewCustomCache()

    for i := 0; i < 5; i++ {
        // Create a key, value, and TTL for each item.
        key := []byte("key" + strconv.Itoa(i))   
        value := []byte("value" + strconv.Itoa(i))
        // Increment TTL by 2 seconds for each item and set it in the cache.
        ttl := time.Duration((i+1)*2) * time.Second 
        c.Set(key, value, ttl) 
        fmt.Printf("Set %s = %s with TTL %v\n", key, value, ttl)
    }

    // Check the cache items at 1-second intervals.
    tick := time.NewTicker(1 * time.Second)
    defer tick.Stop()

    // Loop over the ticker's channel; runs until all items have expired.
    for range tick.C {
        fmt.Println("\nChecking cache status...")
        allExpired := true // check expiration of all items.

        // Check each item in the cache.
        for i := 0; i < 5; i++ {
            key := []byte("key" + strconv.Itoa(i))
            value, ttlLeft := c.Get(key) // Retrieve the item from the cache.
            if value != nil {
                allExpired = false // If the item exists, reset the flag.
                fmt.Printf("Key: %s, Value: %s, TTL Left: %v\n", key, value, ttlLeft)
            } else {
                fmt.Printf("Key: %s has expired\n", key)
            }
        }

        // Break out of the loop once all items have expired.
        if allExpired {
            fmt.Println("All items have expired, stopping...")
            break
        }
    }
}
