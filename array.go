package main

import "fmt"

func main(){
	// zero value array of len 5
	var a [5]int
	fmt.Println("emp:", a)

	// get and set index values
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))


	// one line declare and value definitions
	b := [5]int{2, 4, 5, 29,7}
	fmt.Println("dcl:", b)

	// compiler count the size from the declaration
	b = [...]int{4, 6, 7, 3, 72}
	fmt.Println("size:", len(b))


	// two dimensional 2D Array - type compostion

	var twoD [2][3]int

	for i := range 2{
		for j := range 3{
			twoD[i][j] = i + j
		}
	} 

	fmt.Println("2d:", twoD)

// twoD one line declare
	twoD = [2][3]int{
		{3, 6, 2},
		{5, 8, 2},
	 }

	fmt.Println("2d:", twoD)

}