package main

import "fmt"

func main(){
	// Iterating over slices
	s := []int{2, 4, 5}
	total := 0

	for _, num := range s{
		total += num
	}
	fmt.Println("total:", total)


	// Iterating over maps

	m := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range m{
		fmt.Printf("%s -> %s\n", k,v )
	}


	// Iterating over string characters
	 for i,c :=  range "go"{
		fmt.Println(i, c)

		fmt.Println(i, string(c))
	 }

}