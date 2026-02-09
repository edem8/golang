package main

// The primary mechanism for managing state in Go is communication over channels
// Atomic countere is another option


import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	// atomic integer type- always postive counter.
	var ops atomic.Unit64

	// wait for all rountine to finish work
	var wg sync.WaitGroup

	// each routine adds 1, 1000 times
	for range 50 {
		wg.Go(func(){
			for range 1000{
				ops.Add(1)
			}
		})
	}

	wg.Wait()
	// Load()- saftey to automatically read a value even while other routines are wirting to it
	fmt.Println("ops:", ops.Load())



	// note that using a non atomic integer and then a ++ on it would result in a different number betweenruns- 
	// routines would interfere with each other. And would also produce  race failures with -race flag
}