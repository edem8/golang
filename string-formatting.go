package main
import (
	"fmt"
	"os"
)


type point struct{
	x, y int
}

// alias the print
var pr = fmt.Printf

func main(){

	p := point{1, 3}

	// printing an instance of our struct
	pr("struct1: %v\n", p)

	// This variant of verb includes the fields names
	pr("struct2: %+v\n", p)

	// This prints the Go representation of the value. ie the source code
	// that will produce that value
	pr("struct1: %#v\n", p)






	// To print the type of a value, use %T
	pr("type: %T\n", p)

	// boolean formatting
	pr("bool: %t\n", true)

	// formatting integers- many format however use %d for standard
	pr("int: %d\n", 123)


	// binary formats
	pr("bin: %b\n", 14)
 
	// character formats- print character correspoding to the given integer
	pr("char: %c\n", 33)

	// hex encoding
	pr("hex: %x\n", 456)

	// float format- many as well but standard %f
	pr("float1: %f\n", 79.3)
	pr("float2: %e\n", 123400000.0)
	pr("float3: %E\n", 123400000.0)


	// string formats
	pr("str1: %s\n", "\"string\"")
	pr("str2: %q\n", "\"string\"") //double quote string
	pr("str3: %x\n", "hex this") //renders string in base-16 with two output character per byte


	//print pointer representation
	pr("pointer: %p\n", &p)


	// control width and precision of the rsult when formatting numbers
	pr("width1: |%6d|%6d|\n", 12, 345) //padds with spaces and justifies to the right
	pr("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	pr("width2: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // - minus sign justifies left
	pr("width4: |%6s|%6s|\n", "foo", "b")
	pr("width5: |%-6s|%-6s|\n", "foo", "b")


	// unlike Printf, Sprintf- doesnt print the formatted string to os.stdout.
	// Sprintf return the formatted string without printing

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)


	// Format + print to io.Writers using Fprintf
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	












}