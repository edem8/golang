package main

import "fmt"

// pointers allow for passing references to values and records in a program

func zeroval(ival int){
	fmt.Println(ival)
	ival = 0 

}

func zeroptr(iptr *int){
	*iptr = 0


}


func main(){

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)


}