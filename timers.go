// timers- executing Go code at some point in the future
package main

import (
	"fmt"
	"time"

)

func main(){
	// provides and CHANNEL after the wait time
	timer1 := time.NewTimer( 2 * time.Second)


	<-timer1.C
	fmt.Println("Timer 1 fired")


	// stopping the timer before the operation completes
	timer2 := time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("Timer 2 fired")

	}()
	
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}


	time.Sleep(2 * time.Second)
}