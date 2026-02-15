// Environment variables are unoversal mechanism for conveying configuration information
// to Unix programs. 

package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	// TO set a key/value pai, we use os.Setenv
	// To get a value for a key, we use os.Getenv
	// This will return an empty string if the key isnt present in the environment

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// We can use os.Environ to list all key/value pairs in the environment.
	// It returns a slice of string inthe form KEY=value.
	// we can use strings.SplitN to separate the key and value

	fmt.Println()
	for _, e := range os.Environ(){
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	
	}
}