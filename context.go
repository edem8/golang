// A context carries deadlines, cancellation signals, and other request-scoped values
// across API boundaries and goroutines

package main

import(
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request){
	
	// context is creates for each request by the net/http machinery and
	// is available via the Context method
	
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Println(w, "hello\n")
	case <- ctx.Done():

		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}


func main(){
	// register the handler and start serving
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}