package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				atomic.AddInt32(&counter, 1)
			}()
		}
	}
	wg.Wait()
	fmt.Println("Counter:=", counter)
}
