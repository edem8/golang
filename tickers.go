// timers- when you want to do something once in the future
// tickers- when you want to do something repeatedly at regualr intervals

package main

import (
	"fmt"
	"time"
)

func main(){

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	// just as timers, a channel is sent values
	// using a select to await the values as they arrive

	// routine that executes indepedently of the main
	go func(){
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at:", t)
			}
		}
	}()

	// sync with routine- keep alive long enough for go to fire
	time.Sleep(1600 * time.Millisecond)
	done <- true
	fmt.Println("Ticker program stopped")
}


