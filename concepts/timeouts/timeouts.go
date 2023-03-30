package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string, 1)

	// immediately executed goroutine that sleeps for 2 seconds and then pushes a message to the channel
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "result 1"
	}()

	select {
	// If we receive a response from the channel first
	case res := <-channel1:
		fmt.Printf("Received response from channel 1: %s\n", res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout on channel 1")
	}

	// immediately executed goroutine that sleeps for 2 seconds and then pushes a message to the channel
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "result 2"
	}()
	select {
	case res := <-channel1:
		fmt.Printf("Received response from channel 1: %s\n", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout on channel 1")
	}
}
