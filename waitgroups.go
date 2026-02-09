// To wait for multiple go routines
// Tracks how many routines are running - like a counter
// the w.Wait() in main keeps the main alive until the counter reached 0

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int){
	fmt.Println("Worker", id, "starting")
	time.Sleep(time.Second)

	fmt.Println("worker", id, "done")
}


func main(){
	var wg sync.WaitGroup

	// note: if a wait group is explicitly paassed into functions, it should be done by pointer


	for i:=1; i <=5 ; i++{
		wg.Go(func(){
			worker(i)
		})
	}

	wg.Wait()
}