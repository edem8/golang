package main

import (
	"fmt"
	"errors"
)

type argError struct{
	arg int
	message string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int)(int, error){
	if arg == 42{
		return -1 , &argError{arg, "Can't work with it"}
	}
	return arg +3, nil
}


func main(){
	_, err := f(42)

	var ae *argError 

	// error.As is an advance versio of error.Is. It checks
	// that a given error, matches a specific error type and converts to a
	// value of that type, returning true. and false if there's no match
	if errors.As(err, &ae){
		fmt.Println(ae.arg)
		fmt.Printf(ae.message)
	}else{
		fmt.Println("err doeesnt match argError")
	}
}