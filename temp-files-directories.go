// Throug the program execution we often want to create data that isnt
// needed after the program exists
// Temporary file and directories are useful for this purpose

package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func check(e error){
	if e != nil{
		panic(e)
	}
}


func main(){
	// The easiest way to create a temp file is by calling os.CreateTemp.
	// It creates a file and opens it for reading and writing.
	// We provide "" as the first argument so os,CreateTemp will create it in the default temp directory and a prefix for the filename as "sample"

	f, err := os.CreateTemp("", "sample")
	check(err)

	// The file name begins with the prefix passed as a second argument to CreateTemp. The rest
	// is automatically chosen to ensure concurrent calls will always create different file names. The generated file name is printed to the standard output.
	fmt.Println("Temp file name:", f.Name())


	defer os.Remove(f.Name()) // explicit clean up although os is likely to clean up temp files by itself

	// writing to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// If we intend to write many temp files, we may prefer to create a temporary directory.
	// os.MkdirTemp's arguments are the same as os.CreateTemp, but it creates a directory instead of a file and returns its name.

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname) // clean up the directory after we are done with it

	// synthesize tempory files by prefixing them with our temp directory 
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2, 3, 4}, 0644)
	check(err)

}