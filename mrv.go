package main


import "fmt"

// (int, int) - multi return types

func vals()(int, int){
	return 4, 7
}



func main(){
// often useful for idomatic go
// e.g return both error and result from a func

 a, b := vals()
 fmt.Println(a)
 fmt.Println(b)


//  return subset-some of the values use the _ identifier
_, c := vals()
fmt.Println(c)
 
}