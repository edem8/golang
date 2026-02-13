package main

import (
	"fmt"
	"net/url"
	"net"
)

func main(){
	// sample url for parsing
	s := "postgres://user:pass@host.com:5432/path?k=v#f"


	// Parisng the url and handling err
	u, e := url.Parse(s)
	if e != nil{
		panic(e)
	}

	// Accesing various parts of the url
	// scheme
	fmt.Println(u.Scheme)

	// User. Also user contains all authentication info
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The Host. Also the Host contains the port name and number
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Path
	fmt.Println(u.Path)

	// Fragment
	fmt.Println(u.Fragment)


	// Query params.
	// Get the parameter in a raw string format
	fmt.Println(u.RawQuery)

	// Get the parameter in a map of string to slices of strings
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)

	// can access the vlaues using indexing
	fmt.Println(m["k"][0])
}