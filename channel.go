package main

// channels - pipes that connect concurrent goroutines
// can send values from one go routine and recieve them in another go routine

import "fmt"

func main(){
	// create a channel
	// channels are typed by the values they convey
	messages := make(chan string)


	// sending a "ping" to message
	go func(){
		messages <- "ping"
	}()

	// recieve the "ping" message from above and print
	msg := <-messages
	fmt.Println(msg)
}


// By defualt send and recieve blocks until sender and reciever are both ready hence we dont have to
// use any other synchronizations
