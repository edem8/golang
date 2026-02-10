// panic typically mean something weent unexpectedly
// mostly used to fail dast on errors that shouldnt occur during normal operation
// or errors that we arent prepared to handle gracefully.


package main

import (
	"os"
	"path/filepath"
)


func main(){
	// fails fast and loud and the rest of the program doesnt execute.
	panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil{
		panic(err)
	}
}