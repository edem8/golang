// The go slices package implements sorting for builtin types
// and user defined types

// sortings func is generic and work for any ordered builtin types- int, string etc
// see cmp.ordered for a list of types
package main
import (
	"fmt"
	"slices"
)

func main(){
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)


	ints := []int{5, 2, 7}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	// check if slice is sorted
	s := slices.IsSorted(ints)
	fmt.Println("sorted:", s)
}

