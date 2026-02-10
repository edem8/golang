package main



import (
	"bytes"
	"fmt"
	"regexp"
)

func main(){
	// test whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+ch)", "peach")
	fmt.Println(match)

	// Above, we use a string pattern directly, byfor other regExp tasks
	// we will need to compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")


	// The struct has many methods

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch")) //finds a match for the regEx
	fmt.Println("idex:", r.FindStringIndex("peach punch")) //find a match and return the start and end index
	fmt.Println(r.FindStringSubmatch("peach punch")) //return info about both the full match and submatch
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))


	fmt.Println(r.FindAllString("peach punch pinch", -1)) // find matched in all string not just the first
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //all match applicable to other funcs
	fmt.Println(r.FindAllString("peach punch pinch", 2))  //non- negative integer will limit the number of matches
	
	
	// Instead of string arguments we can
	fmt.Println(r.Match([]byte("peach")))


	// when creating global regex, its safer to use MustCompile instead of Compile
	// Since the former return a panic instead of an error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "fruit"))


	// The function variant helps to transform matched text with a function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}