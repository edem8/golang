// sha256 hases are used to compute short identities for binaries 
// or text blobs

package main

import (
	"fmt"
	"crypto/sha256"
)


func main(){
	s := "sha256 is a string"

	// create a new hash
	h := sha256.New()

	// write accepts bytes. If you have a string 
	// use []bytes(str) to convert it to bytes

	h.Write([]byte(s))


	// finalize the hash result as a slice of bytes
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs)
}