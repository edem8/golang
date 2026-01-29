package main

import "fmt"

func main(){

	// basic
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

// initial, condition, after
	for j:=0; j<9; j++{
		fmt.Println(j)
	}

	// N-times | range
	for i:= range 3{
		fmt.Println("range", i)
	}

	// Infinte loop until break
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n % 2 == 0{
			continue
		}
		fmt.Println(n)
	} 
}