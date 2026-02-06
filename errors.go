package main
import (
	"errors"
	"fmt"
)

// By convention errors are returned as the last value 
func f(arg int) (int, error){
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	//nil in place of error mean that there was no error
	return arg + 3, nil
}

// sentinel error- predeclared variable used to signify a specific error condition
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")


func makeTea(arg int) error {
	if arg == 2{
		return ErrOutOfTea
	}else if arg==4{
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}


func main(){
	for _, v := range []int{7, 42}{
		if r, e := f(v); e != nil{
			fmt.Println("f failed:", e)

		}else{
			fmt.Println("f worked:", r)
		}
	}




	for num := range 5{
		if err := makeTea(num); err != nil{
			if errors.Is(err, ErrOutOfTea){
				fmt.Println("We should buy more tea")
			}else if errors.Is(err, ErrPower){
				fmt.Println("Now it is dark")
			}else{
				fmt.Println("unkown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready")
	}


}