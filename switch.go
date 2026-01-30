package main

import (
	"fmt"
	"time"
)



func main(){
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1: 
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3: 
		fmt.Println("three")
	}



	// comma separated multiple switch expressions in same case
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's a Weekend")
	
	default:
		fmt.Println("It's a Weekday")

	}


	// switch with no expression is an alternate if-else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("it's afternoon")
	}


	// Type switch - comparing types instead of values
	whoami := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("I am boolean")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("I am %T\n", t)
		}
	}

	whoami(true)
	whoami(3)
	whoami("chale")

}