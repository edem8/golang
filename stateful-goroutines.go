// state management with goroutines and channels
// This channel based approach aligns with Go's idea of sharing memory 
// by commuincating and ahving each piece of data owned by exaclty  1 goroutine
package main


import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// prevent concurrent access - 1 routine ownership
// read and writes are via channel send and recieves
 

type readOp struct{
	key int
	resp chan int
}

type writeOp struct{
	key int
	val int
	resp chan bool
}

func main(){
	// tracking the number of read and write operations
	var readOps uint64
	var writeOps uint64

	// These channnels will be used by other routines to issue read/write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// state owner ...unlike the MUTEX approach the state (also a map) is 
	// more private to this routine
	go func(){
		var state = make(map[int]int)
		for {
			select{
			case read := <-reads:
				read.resp <-state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp  <- true
			}
		}
	}()


	// read go routine -issues 100 requests
	for range 100{
		go func(){
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
	
	
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
	
			}
		}()
	}


	// write routine - issues 10 writes
	for range 10{
		go func(){
			for{
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
	
				writes <- write
				<-write.resp
	
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
	
			}
		}()
	}



	// note that the loops in the routines run infinitely - but the program is only kept alive 
	// for a second only
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)


	writeFinalOps := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeFinalOps)


}
