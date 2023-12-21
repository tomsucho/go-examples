package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Closing!")
				return
			case t := <-ticker.C:
				fmt.Println("Ticket at:", t)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	done <- true
}
