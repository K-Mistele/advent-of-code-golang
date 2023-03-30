package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a channel of requests, with a buffer of 5
	requests := make(chan int, 5)

	// push 5 requests into the channel
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Receive a value every 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	// Receive a request every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Create a bursty limiter with a capacity of 3
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel; simulating an allowed burst
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 milliseconds, push a request into the channel
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Create a channel of requests (5) and fill it
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
