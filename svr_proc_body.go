

package main

import (
	"fmt"
	"io/ioutil"
	"log"  //~ What the fuck is log....(from library??)
	"net/http"
	"net/url"
)

//~ Do not go copy.
//~ The type of req.Form : map[string]string
func printMap(form url.Values) {
	fmt.Println("BEGIN<PRINTING>")
	for key, value := range(form) {
		fmt.Printf("<%v> ->  <%v>\n", key, value)
	}
	fmt.Println("END<PRINTING>")
}

func printBody(req *http.Request) {
  defer req.Body.Close()
	inBuff, err := ioutil.ReadAll(req.Body)
	if nil != err {
		fmt.Println("Error reading body of request")
		panic(err)
	}
	fmt.Printf("Body length is :%v\n", len(inBuff))
	fmt.Printf("Body is:<%v>\n", string(inBuff))
}

func main() {
	//~ And this shall be in a separate go-routine. 
	mux := http.NewServeMux()

	// mux.Handle("/api/", func(w http.ResponseWriter, req *http.Request){
	// fmt.Println("Access to /api/")
	// fmt.Fprintf(w, "api, not founded yet")
	//})
	//mux.Handle("/api/", apiHandler{})

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		//~ OK. I need to fetch the post data.
		printBody(req)
		fmt.Println("Request in")
		if req.URL.Path != "/" {
			fmt.Println("Accessing to <" + req.URL.Path + ">")
			http.NotFound(w, req)
			return
		}
		fmt.Println("Yes")
		fmt.Fprintf(w, "homepage")
	})
	//~ The second parameter is a `MUX'
	log.Fatal(http.ListenAndServe(":12345", mux))
}
	
