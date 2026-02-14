// The filepath package provides functions to parse and construct file paths
// in a wasy that is portable between operating systems; dir/file on linux and dir\file on windows

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main(){
	// Join should be used to create a file path in a portable way
	// it takes any number of arguments and constructs a hierarchical path from them
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	// Always use join instead of concatenating /s or \s manually. In addition to provding
	// portability, Join will also normalize paths by removing superfluous separators and directory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// The Dir and Base can be used to split the a path to the directory and the file.
	// Alternatively, SPlit will return both in the same call.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))


	filename := "config.json"

	// some files names have extensions following a dot.
	// We can split the extensions out of such names with Ext.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// to find the filename with extension removed use strings.TrimSuffix
	fmt.Println(strings.TrimSuffix(filename, ext))


	// rel finds a relative path between a base and a target.
	// returns an error if the target cannot be made relative to the base
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil{
		panic(err)
	}
	fmt.Println(rel)


	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil{
		panic(err)
	}

	fmt.Println(rel)

}