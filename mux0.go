

package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{
}

//~ interface of http.Handler:
func (p*MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/":
		sayHello(w, r)
	default:
		http.NotFound(w, r)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, my route")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
