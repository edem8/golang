package main

import "fmt"



// closures = anonymous function (an inline function you dont want to name)

func intSeq() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}

func main(){
	
	// returns a function to nextInt with its own state of i
	nextInt := intSeq()


	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	// func2 with its own state
	newInt := intSeq()
	fmt.Println(newInt())



}