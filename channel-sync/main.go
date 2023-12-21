package main

import "fmt"

func main() {

	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			if j, more := <-jobs; more {
				fmt.Println("Job handler: received", j)
			} else {
				fmt.Println("Job handler: channel closed..")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Sending job", i)
		jobs <- i
	}
	close(jobs)
	<-done

}
