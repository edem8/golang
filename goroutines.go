// A lightweight threas of execution
package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i := range 3 {
		fmt.Println(from, ":", i)
	}

}

func main(){
	// calling the function synchronously
	f("direct")


	// invoking the function in a go routine
	go f("go routine")

	// routine for anonymous func call
	go func(msg string){
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

