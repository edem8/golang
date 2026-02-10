// Recover can stop a panic from aborting the program and let it continue
// with execution instead.


// usecase example- a sercer wouldnt want to crash if one of the client connections exhibit a critical
// error. INstetad the server would want to close that connect and server other clients

package main

import "fmt"


func mayPanic(){
	panic("a problem")
}

func main(){
	// recover must be called within a deferred function.
	// When the enclosing func panics, a recover call within it will catch the panic
	//  The return value of the recover is the error raised in the panic

	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// The code below will not run because mayPanic panics. The execution of
	// main stops at this points and resumes in the deferred closure
	fmt.Println("After the panic")
}