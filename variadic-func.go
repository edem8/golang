package main


import "fmt"

// funcs that can be called with any numbeof trailing parameter eg. fmt.println()
func sum(nums ...int){

	// within the func, num is equivalent to []int - a slice
	// As such you can perform equivalent operation on it
	fmt.Print(nums)

	total := 0
	for _, num := range nums{
		total += num
	}

	fmt.Println(total)



}



func main(){
    sum(1, 2, 3)

	// Passing a slice to a variadic function
	nums := []int{3, 5, 6}
	sum(nums...)
}