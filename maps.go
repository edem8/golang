package main

import (
	"fmt"
	"maps"
)


func main(){
	// maps= hashes = dict (associative data types)
	// map[key-type]value-type

	m := make(map[string]int) 
	fmt.Println("emp:", m)


	// get and set map values
	m["k1"] = 7
	m["k2"] = 12
	fmt.Println("set:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))
	

	// delete vlaue from a map
	delete(m, "k1")
	fmt.Println("del:", m)

	// clear the map - remove all data
	clear(m)
	fmt.Println("clr:", m)


	// optional check to make sure key was present when retrieving a value
	// _ because we dont need the value just present checks
	// useful for differentiating between non present keys and keys with empty or 0 values e.g "" or 0 
	_, present := m["k1"]
	fmt.Println("Prs:", present)



	// one line definition of maps
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// maps package utility funcs
	if maps.Equal(n, n2){
		fmt.Println("n == n2")
	}
}