package main

import (
	"fmt"
	"slices"
)


func main(){
	// slices-typed by the elements they contain
	// uninitialized slices have 0 len and equals nil( nothing inside)

	var s []string
	fmt.Println("uninit:", s, s == nil, "len:", len(s))


	// non-zero slices- defined length and capacity(when we now the slices will grow overtime)
	s = make([]string, 3, 8)
	fmt.Println("empt:", s, "len:", len(s), "cap:", cap(s))



	// get and set slices values
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))


	// slices support append operations and return a slices
	s = append(s, "d")
	s = append(s, "e", "f") 
	fmt.Println("apd:", s)


	// slices support copy operations
	c := make([]string, len(s))
	copy(c,s)

	fmt.Println("copy:", c)


	// slices support "slice" operations
	l := s[2:5]
	fmt.Println("sl1:", l)	

	l = s[:5]
	fmt.Println("sl2:", l)
	
	l = s[2:]
	fmt.Println("sl3:", l)	

	
	// one line slices declaration
	t := []string{"4", "6", "7", "8", "2"}
	fmt.Println("dcl:", t)	

	// slices pacakage has utils funcs
	t2 := []string{"4", "6", "7", "8", "2"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2")
	}



	// Just as we can make 2D arrays, we can make 2D slices but the lenghth of inner slices can vary
	// [[0] [1,2] [2,3,4]]

	twoD := make([][]int, 3)
	for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}