

package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request){
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":8081", nil)
}
