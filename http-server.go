// writing a basic HTTP server using net/http package

package main

import (
	"fmt"
	"net/http"
)


// A fundamental concept in net/http servers is handlers
// A handler is an object that implements the http.Handler interface
// A common way to write a handler is by usign the http.HandleFunc adapter on functions
// with appropriate signature

func hello(w http.ResponseWriter, req *http.Request) {


	// functions serving as handler take a http.RequestWriter and a http.Request as arguments
	// The response writer is used to fill the http reesponse. Here the response
	// is simply Hello\n
	fmt.Fprintf(w, "hello\n")
}


// Handlers do something alittle more sophisticated by reading all http request headers and echoing them into the response body
func header(w http.ResponseWriter, req *http.Request){
	for  name, headers := range req.Header{
		for _, h := range headers{
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}


func main(){
	// We register our handlers on server routes using the http.HandleFunc
	// convinience function. Itsets up the defualt router in the net.http packege and takes a
	// fucntion as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", header)


	// Finally, we call the listen and server with the port and a handler. nil tellls
	// it to use the defualt router we've just setup
	http.ListenAndServe(":8090", nil)
}