package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Imaginary worker function
func someWorker(id int, messages chan int) {
	timeToSleep := rand.Intn(10)
	fmt.Printf("Worker %d sleeping for %d seconds\n", id, timeToSleep)
	time.Sleep(time.Second * time.Duration(timeToSleep))
	messages <- id
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	timesToSleep := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go someWorker(i, timesToSleep)
	}

	var message int
	for i := 0; i < 10; i++ {
		message = <-timesToSleep
		fmt.Printf("Worker %d done sleeping\n", message)
	}

	time.Sleep(time.Second * time.Duration(12))
}
