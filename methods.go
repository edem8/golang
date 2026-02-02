package  main

import "fmt"

// methods - functions with a special receiver argument on struct types
type rect struct{
	width, height int
}

func (r *rect) area() int{
return r.width * r.height
}

func (r rect) perimeter() int{
	return 2 * (r.width + r.height)
}


func main(){
	 r := rect{width: 10, height: 5}

	 fmt.Println("area:", r.area() )
	 fmt.Println("perimeter:", r.perimeter())

	 sr := &r
	 fmt.Println("area:", sr.area() )
	 fmt.Println("perimeter:", sr.perimeter())
}