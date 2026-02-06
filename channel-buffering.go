package main

// By defualt are unbuffered-they will only accept(chan <-) if
// there is a corresponding recieve(<-chan) ready to accept the sent value

// Buffered channels accepts a limited number of values without a corresponding reciever

import "fmt"

func main(){

	// make a channel buferring 2 values
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}