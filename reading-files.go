package main

import (
	"fmt"
	"path/filepath"
	"os"
	"io"
)


// Helper func to check for errors
func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	path := filepath.Join(os.TempDir(), "dat")

	// Reading the entire file into memory
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Println(string(dat))



	// controlling how and what parts of the file are read.
	// open the file for reading
	f, err := os.Open(path)
	check(err)

	// read the first 5 bytes of the file into b1
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))


	// Seeking to a known location in the file and reading from there
	o2, err := f.Seek(6, io.SeekStart) //return the current byte position and an error
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2) //read rreturn the number of bytes read and an error after  reading into memory(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	fmt.Printf("%v\n", string(b2[:n2]))

	// other methods of seeking relative to the cursor
	// o3, err := f.Seek(2, io.SeekCurrent)
	// check(err)

	// o4, err := f.Seek(-4, io.SeekEnd)
	// check(err)

	// the io package provides helpful functions for file reading.
	// for example, reads like the above can be performed with more robust functions
	// ReadAtLeast

	o5, err := f.Seek(6, io.SeekStart)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o5, string(b3))

	// there is no built in rewind-way to go back but seek can do that still
	_, err := io.Seek(0, io.SeekStart)
	check(err)

	// The bufio package implements a buffered reader that may be useful both for its efficiency
	// with small reads and addittional reading methods
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) // returns the next n bytes without advancing the reader
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.close()

}