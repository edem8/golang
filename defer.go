// defer - used to enusre that a function call is performed later
// in a program's execution, usally for
// purposes of cleanup

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	path := filepath.Join(os.TempDir(), "defer.text")
	f := createFile(path)
	// executed at the end of the enclosing func - in this case main
	defer closeFile(f)
	writeFile(f)
}


func createFile(p string) *os.File{
	fmt.Println("creating...")
	f, err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}


func writeFile(f *os.File){
	fmt.Println("Writing...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File){
	fmt.Println("closing...")
	err := f.Close()
	if err != nil{
		panic(err)
	}
}