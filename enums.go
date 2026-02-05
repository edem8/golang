package main


import "fmt"

// implementing enums using existing language idoms - go doesnt have enum type
// as a distinct lang feature
type ServerState int 


// enum-like pattern 0,1,2,...
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",

}


// stringer interface-printing a state calls the string method automatically
// think of this like a custom string print without it, it will print the assigned value
// when you print a server state in this case 0, 1, 2...
func (ss ServerState) String() string{
	return stateName[ss]
}


func main(){

// example print-automatic call
fmt.Println(StateIdle)

}

