package main

import (
	"math"
	"fmt"
)

// const - declares a vlaue as const
const s string = "constant"

func main(){
    fmt.Println(s)

	// constants can be defined within a func scope
	const n = 500000000

	// constant expression performs arithmetic with high precision
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it is given one, such as by an explicit conversion
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or a function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}