package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/fs"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	// creae a new sub-directory in the current working directory
	err := os.Mkdir("subdir", 0755)
	check(err)

	// when creating temp directories, it is good practice to defer their removal them after
	// similar to rm -rf
	defer os.RemoveAll("subdir")


	// Helper functions to create a new empty file
	createEmptyFile := func(name string){
		d := []byte("")
		check(os.WriteFile(name, d, 0644))

	}
	createEmptyFile("subdir/file1")

	// create hierarchical directories with MkdirAll similar to mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/child/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/file4")

	// ReadDir list directory contents, returning a slice of os.DirEntry objects
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("listing subdir/parent:")
	for _, entry := range c{
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// chdir lets you change the current working directory similar to cd
	err = os.Chdir("subdir/parent/child")
	check(err)


	// Now if we list the directory again we will see the content of the child directory
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("listing subdir/parent/child:")
	for _, entry := range c{
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started
	err = os.Chdir("../../..")
	check(err)

	// we can visit a directory recursively, including all its sub-directories.
	// WalkDir accepts a callback function to handle every file or directory visted
	fmt.Println("visiting subdir:")
	err = filepath.WalkDir("subdir", visit)



	
}

func visit(path string, d fs.DirEntry, err error) error{
	check(err)

	fmt.Printf("visited: %s (isDir: %t)\n", path, d.IsDir())
	return nil
}