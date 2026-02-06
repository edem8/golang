// typed parameters (generics)

package main

import(
	"fmt"
	// "cmp"
)


// type Ordered interface{
// 	int | string | float64
// }

//suing the package instead of interface
// func max[T cmp.Ordered](a T, b T) T{
// 	if a > b{
// 		return a
// 	}
// 	return b
// }

// func main(){
// fmt.Println(max(1,2))
// fmt.Println(max("hello", "world"))
// fmt.Println(max(1.4, 2.5))

// }





// func Index[T comparable](list []T, target T) int{
// 	for idx, value := range list{
// 		if value ==target{
// 			return idx
// 		}
// }
// return -1

// }


// func main(){
// 	listInt :=  []int{1, -20, 40}
// 	fmt.Println(Index(listInt, -20))
// }



type Box[T any] struct{
	items []T
}

func (b *Box[T]) Add(item T){
b.items = append(b.items, item)
}

func (b *Box[T]) Get() []T{
	return b.items
}


func main(){
	b := Box[int]{}
	b.Add(10)
	b.Add(20)
	b.Add(40)
	// b.Add("hello")

	for _, item := range b.Get(){
		fmt.Println(item * 10)
	}

	// for _, v := range b.Get() {
	// 	fmt.Printf("value=%v, dynamic type=%T\n", v, v)
	// }
	
}