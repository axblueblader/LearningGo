package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("demo")
	maxTime := 10
	timeout := time.After(time.Duration(maxTime) * time.Second)
	requestChan := make(chan int)
	data := 0
	go func() {
		for {
			time.Tick(2 * time.Second)
			data++
			requestChan <- data
		}
	}()
	for {
		select {
		case <-timeout:
			fmt.Println("Server time out")
			return
		case request := <-requestChan:
			go doRequest(request)
		default:
			waitingRequest()
		}
	}
}

func doRequest(request int) {
	fmt.Printf("Request recieved and being processed: %d\n", request)
}

func waitingRequest() {
	fmt.Println("Waiting for request")
	time.Sleep(1 * time.Second)
}
