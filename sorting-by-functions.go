// sometimes we want to sort a collection by something
// other than its natural order. For example

package main

import (
	"fmt"
	"slices"
	"cmp"
)


func main(){
	fruits := []string{"orange", "banana", "kiwi"}


	// custom func to sort based on length
	lenCmp := func(a, b string) int {	
		return cmp.Compare(len(a), len(b))

	}	

	// SortFunc
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)



	// Implementation 2- on custom types (non built in type)
	type Person struct{
		name string
		age int
	}


	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}


	slices.SortFunc(people, func(a, b Person) int{
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)
}