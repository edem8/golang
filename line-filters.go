// A line filter is a program that reads input
// from a stdin and writes to stdout. grep and sed are common line filters

package main


import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	// Example that writes a capitalized version of all input text.
	//wrapping the unbuffered stdin with a buffered scanner, give a convinient scan mathod that advances the scanner
	// to the next token-which is the next line in the defualt scanner
	scanner := bufio.NewScanner(os.Stdin)


	for scanner.Scan(){
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// check if errors during scan.
	// End  of file is expected and not reported as an error

	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}