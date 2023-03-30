package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Send a ping
func ping(id int, pings chan<- string) {
	// sleep for random amount of seconds
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	pings <- fmt.Sprintf("ping from %d", id)
	return
}

//  pong in response to a ping
func pong(pings <-chan string, pongs chan<- string) {
	pingMessage := <-pings
	fmt.Printf("received ping (%s)\n", pingMessage)
	pongs <- fmt.Sprintf("pong ACKing ping (%s)", pingMessage)
}
func main() {

	pings := make(chan string, 10)
	pongs := make(chan string, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("Dispatching pinger %d\n", i)
		go ping(i, pings)
		go pong(pings, pongs)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("received pong (%s)\n", <-pongs)
	}

	// mess with closing a channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	// immediately executed goroutine
	go func() {

		// Infinite loop to process jobs until we're done
		for {
			j, more := <-jobs // more is true IFF there are more jobs OR the channel is not closed
			if more {
				fmt.Println("received a job", j)
			} else {
				fmt.Println("All jobs received")
				done <- true
				return
			}
		}

	}()
	for j := 1; j <= 3; j++ {
		jobs <- j // push 3 jobs to the channel
		fmt.Println("Send job", j)
	}
	close(jobs) // close the channel
	fmt.Println("Send all jobs and closed the channel")

	<-done // wait for goroutine to finish
}
