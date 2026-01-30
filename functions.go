package main

import "fmt"

// returns the sum of two int

func plus(a int, b int) int {
return a + b
}


// mutliple params of the same tpye, you can omit the types
// until the last like-type parameter

func plusplus(a, b, c int) int {
	return a + b + c
}



func main(){
// function calls
sum1 := plus(4, 5)
fmt.Println("sum1:", sum1)


sum2 := plusplus(4, 5, 2)
fmt.Println("sum2:", sum2)
}