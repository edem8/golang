package main
import (
	"fmt"
	"unicode/utf8"
)


func main(){

	// strings are []bytes -slice of bytes
	const s = "สวัสดี"

	fmt.Println("len:", len(s)) //prints the len of bytes stored


	// looping through the strings indexes into the raw byted produced
	for i :=0 ; i < len(s); i++{
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	//characters are called runes in go- integer representation of unicode points
	//utf8 package has utlis for peforming rune operations

	// Rune count
	fmt.Println("Runce Count:", utf8.RuneCountInString(s))


	// indexing and decoding rune
	// The range loop indexes and decodes runes more easily

	for idx, runeValue := range s{
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println()
	fmt.Println("\nUsing DecodeRuneInString")
	// achieving the above with for loop requirement a util
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}


}