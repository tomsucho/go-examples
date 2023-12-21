package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan string) {
	for job := range jobs {
		fmt.Printf("Worker id:=%d starting job:=%d\n", id, job)
		time.Sleep(time.Second)
		results <- fmt.Sprintf("Worker id:=%d finished job:=%d\n", id, job)
	}
}

func main() {

	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 0; w < 2; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for k := 0; k < numJobs; k++ {
		fmt.Println(<-results)
	}
}
