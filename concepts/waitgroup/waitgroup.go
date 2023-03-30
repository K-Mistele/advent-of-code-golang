package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) // add a new worker to the waitgroup
		i := i

		go func() {
			defer wg.Done() // call this at the end of the routine to let the waitgroup know I'm done
			worker(i)       // call a synchronous function
		}()
	}

	wg.Wait() // wait for all workers to finish
}
