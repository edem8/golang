package main


import "fmt"

// Embeddding structs and interfaces ti express a more seamless composition of types
type base struct{
	num int
}

//method
func (b base) describe() string{
	return fmt.Sprintf("base with num=%v", b.num)
}

// embedding base in container
type container struct{
	base
	str string
}

func main(){
	co := container{base:base{num: 1,},
str: "some name"}

// direct access
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// alternatively you can spell out the full apth
	fmt.Println("also num:", co.base.num)

	// base methods become container methods
	fmt.Println("describe:", co.describe())

	// embedding also bestows interface implementation on other structs
	type describer interface{
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())


}
