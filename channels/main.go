package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	m := make(chan string, 11)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m <- fmt.Sprintf("Id: %d", i)
			fmt.Println("Sent:", i)
		}(i)
		if i == 9 {
			//time.Sleep(1 * time.Second)
			close(m)
		}

	}
	// wg.Add(1)
	// go func(t time.Duration) {
	// 	defer wg.Done()
	// 	time.Sleep(t * time.Minute)
	// }(1)

	//wg.Wait()
	for {
		if val, ok := <-m; ok {
			fmt.Println("Received:", val)
		} else {
			break
		}
	}
	fmt.Println("Finished!")
}
