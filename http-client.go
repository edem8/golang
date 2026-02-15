// using the net/http package http request

package main

import (
	"fmt"
	"net/http"
	"bufio"

)


func main(){
	// Issue an http GET request to a server
	// http.Get is a convinient shortcut for creating an http.Client
	// object and calling its Get method; It uses the http.DefaultClient object
	// which has useful default settings
	resp, err := http.Get("https://gobyexample.com")
	if err != nil{
		panic(err)
	} 

	defer resp.Body.Close()

	// print the response status
	fmt.Println("Response status:", resp.Status)

	// print the first 5 lines of the response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++{
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil{
		panic(err)
	}
}