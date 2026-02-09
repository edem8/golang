package main

//making a worker pool with goroutines and channels

import (
	"fmt"
	"time"
)

// worker pool
func worker(id int, jobs <-chan int, results chan int){
	for j := range jobs{
		fmt.Println("worker", id, "started job", j)

		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
	

}


func main(){
	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)


	// 3 workers
	// since jobs are empty worker execution will be blocked
	// But notice that three go rountines and running the order they executes depends on the go scheduler
	for w :=1 ; w <=3 ; w++{
		go worker(w, jobs, results)
	}

	// sending jobs
	for j := 1; j <= numOfJobs; j++{
		jobs <- j
	}
	close(jobs)


	// main will stay alive until result have been recieved- technically sent and recieved
	// alternatively we can wait for the multiple go routines to finish using WAITGROUPS especially if we dont need the results
	for a:= 1; a <= numOfJobs; a++{
		<-results

	}
}