package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Container struct {
	mtx   sync.Mutex
	items map[string]int
}

func (c *Container) inc(key string, i int, delay time.Duration) {
	c.mtx.Lock()
	time.Sleep(delay)
	c.items[key] += 1
	defer c.mtx.Unlock()
	fmt.Printf("%d: incremented: %d\n", i, c.items[key])
}

var wg sync.WaitGroup

func main() {
	c := Container{
		items: map[string]int{"a": 0, "b": 0},
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int, t time.Duration) {
			defer wg.Done()
			c.inc("a", i, t)
			//c.inc("b", i)
		}(i, time.Duration(rand.Intn(3)))

	}
	wg.Wait()
	fmt.Println(c.items)
}
