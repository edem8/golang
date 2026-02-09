package main

// RoC - Iterating over channels just as we did over collections and data structures

import "fmt"

func main(){
	queue := make(chan string, 2)
	queue <- "1"
	queue <- "2"

	close(queue)
	
	// ranging- automatic stop when channel is closed
	for elem := range queue{
		fmt.Println(elem)
	}
}
