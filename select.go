package main

import (
	"fmt"
	"time"
)

// go select letss you wait mutiple channel operations
// combining goroutines and channels with select unleashes a powerful go feature

func main(){
	c1 := make(chan string)
	c2 := make(chan string)


	go func(){
		time.Sleep(2 *  time.Second)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()


	for range 2{
		select {
		case msg1 := <-c1:
			fmt.Println("recieved:", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved:", msg2)
		}

		
	}

}