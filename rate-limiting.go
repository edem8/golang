// Rate limiting- mechanism for controlling resource utilization
// and maintaining quality of service.

// implementation using goroutines, channels and tickers

package main

import (
	"fmt"
	"time"
)


func main(){

	//Basic implementation - controlling(handling) incoming requests

	// request channel buffer
	requests := make(chan int, 5)

	for i := 1 ; i <=5; i++{
		requests <- i
	}
	close(requests)


	limiter := time.Tick(200 * time.Millisecond)

	// limiting ourselve to 1 request every 200 millisecond
	for request := range requests{
		<-limiter
		fmt.Println("request", request, time.Now())
	}


	fmt.Println("")
	fmt.Println("")

	// Implementation 2 - allowing short bursts of requests in our scheme
	// basically buffer the limiter channel

	burstyLimiter := make(chan time.Time, 3)

	// fill up channel with current time to burst current (3) reuqest
	for range 3{
		burstyLimiter <- time.Now()
	}

	// fill the channel after every 2sec to to add new value up to its limit
	go func(){
		for t := range time.Tick(200 * time.Millisecond){
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <=5; i++{
		burstyRequest <- i
	}

	close(burstyRequest)

	for request  := range burstyRequest{
		<-burstyLimiter
		fmt.Println("request", request, time.Now())
	}



}
