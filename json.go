// json encoding and decoding from and to builtin and cutom types
package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"os"
)

// structs for custom types encoding and decoding
type response1 struct{
	Page int
	Fruits []string
}

// only exported fields will be encoded /decoded in JSON.
// FIelds must start with capital letters to be exported
type response2 struct{
	Page int	`json:"page"`
	Fruits []string `json:"fruits"`
}


func main(){
	// encoding basic data types to JSON string
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _  := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Encoding slices and maps to JSON arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))


	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))


	// encoding custom data types
	// The JSON package can automatically encode custom types. It will only
	// include exported fields in the encoded output and will by default use
	// those names as JSON keys

	resp1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	resp1B, _ := json.Marshal(resp1D)
	fmt.Println(string(resp1B))

	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "pear", "lemon"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))





	// Decoding JSON into GO values
	// Here's an example for a generic data structure
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	// variable to store the decoded data
	var dat map[string]interface{}


	//decoding the data
	if err := json.Unmarshal(byt, &dat); err != nil{
		panic(err)
	}
	fmt.Println((dat))

	num := dat["num"].(float64)
	fmt.Println(num)


	// accessing the nested data
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)


	
	// Decoding into custom data types
	// This as the advantage of types safety
	// eliminates the need for type assertions when decoding

	str := `{"page": 1, "fruits":["apple","peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Above we used bytes and strings as intermediated for JSON and data represntation
	// we can also stream the encoding directly to os.Writers like Stdout and even HTTP response bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5, "lettuce":7}

	enc.Encode(d)


	// Streaming reads form os.Readers like os.Stdin, and HTTP request bodies is also supported by the json package
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}