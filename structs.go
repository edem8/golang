package main

import "fmt"

// structs - typed colleection of fields(grouping data to form records)
type person struct{
	name string
	age int
} 

func newPerson(name string) *person{

	p := person{name: name}
	p.age = 42
	return &p
}


func main(){
	// create a new struct
	fmt.Println(person{"Bob", 20})

	//named fields
	fmt.Println(person{name: "Alice", age: 30})

	// ommitting fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// using & operator to generate a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// using a function to create new struct
	fmt.Println(newPerson("Jon"))


	// Dot access
    s := person{name: "Sean", age: 50}
	fmt.Println(s.name)


	sp := &s
	fmt.Println(sp.age)


	sp.age = 51
	fmt.Println(sp.age)


	// anonymus struct - usecase is a single value

	dog := struct{
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}