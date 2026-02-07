// timeouts - essential for programs that connect to external resources
// or that otherwise need execution time bounds.


package main

import (
	"fmt"
	"time"
)



func main(){
c1 := make(chan string, 1)

// go routine with non-blocking channel read
// common pattern for preventing goroutine leaks in case chaneel is never read
go func(){
	time.Sleep(2 * time.Second)
	c1 <- "result 1"
}()

// implementing a timeout with select
select{

case result := <-c1:
	fmt.Println(result)
case <-time.After(1 * time.Second):
	fmt.Println("timeout 1")

}


// goroutine 2 - with longer timeout
c2 := make(chan string, 1)
go func(){
	time.Sleep(2 * time.Second)
	c2 <- "result 2"
}()

select{
case result2 := <-c2:
	fmt.Println(result2)
case <-time.After(3 * time.Second):
	fmt.Println("timeout 2")
}

}