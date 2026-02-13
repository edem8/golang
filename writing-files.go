package main

import (
	"fmt"
	"path/filepath"
	"os"
	"bufio"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}


func main(){
	// dumping a string or bytes into a file
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	
	// Opening the file for write more granular writes
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// closing the file but with a defer
	defer f.Close()


	// writing bytes slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)


	// Write strings
	n3, err  := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // flushes the writes to stable storage- the data is not written to disk immediately.. it is put in RAM..This forces the os to actually put it on disk


	// bufio provides buffered writes jsut they do with reads
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // Ensures all buffered operations have been applied to the underlying writer.
}