package main

import "fmt"

// recursive - call itself until a base case is reached

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}


func main(){
fmt.Println(factorial(7))
}