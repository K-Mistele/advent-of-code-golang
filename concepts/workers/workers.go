package main

import (
	"fmt"
	"time"
)

// worker function with a receive-only jobs channel and a send-only results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create workers for 3 jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// push 3 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 1; j < numJobs; j++ {
		<-results
	}
}
