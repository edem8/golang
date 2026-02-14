// the go:embed directive is a compiler directive that allows 
// programs to include arbitary files and folder in th go binary at build time
package main

import (
	"embed"
)



// The embed directive accepts paths relative to the directory containing
// the go source file.
// This directive embeds the contents of the file into the string vriable immediately following it


//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte


// We can also embed multiple files or even folders with widlcards. The variable 
// embed.FS types implements a simple virtual file system.

//go:embed folder/*.hash
//go:embed folder/single_file.txt
var folder embed.FS


func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
