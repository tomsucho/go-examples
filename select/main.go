package main

import (
	"fmt"
	"time"
)

func simulate_work(s string, c chan string, t time.Duration) {
	time.Sleep(t)
	c <- fmt.Sprintf("%s finished!", s)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go simulate_work("c1", c1, 3*time.Second)
	go simulate_work("c2", c2, time.Second)

	//for i := 0; i < 2; i++ {
	select {
	case s := <-c1:
		fmt.Println(s)
	case s := <-c2:
		fmt.Println(s)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
	//}
}
