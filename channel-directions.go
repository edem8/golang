package main
import "fmt"

// when using channels as function params,
// we can specify if a channel is meant to only send or recieve values.

// sending parameter channel
func ping(pings chan<- string, msg string){
	pings<-msg
}


// recieveing parameter channel
func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg

}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}