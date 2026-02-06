package main

import (
	"fmt"
	"time"
)

// channels can be used to synchronize executive across go routines
// eg. below blocking recieve to wait for a goroutine to finish.
func worker(done chan bool){
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// we can send this value even without a reciver because of the buffer
	// aslo noted the rest of the program will wait for this async routine because 
	// until the channel has sent and recieved, the program wont terminate
	done <- true 
}


func main(){
	// buffering one val so we can send even without a reciver ready
	done := make(chan bool, 1)
	
	// go routine
	go worker(done)

	fmt.Println("yes:", <-done)

}