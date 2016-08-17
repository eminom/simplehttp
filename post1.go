
package main

/////////////////////////////////////////////
//////////////// What the fack...........
///////

import (
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"reflect"
)

const url_uno = "http://localhost:12345/"

func printResp(resp *http.Response){
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("StatusCode(type) = ", reflect.TypeOf(resp.StatusCode).Name())
	fmt.Println("Status = ", resp.Status)
	fmt.Println("Proto = ", resp.Proto)
	fmt.Println("ProtoMajor = ", resp.ProtoMajor)
	fmt.Println("ProtoMinor = ", resp.ProtoMinor)
	fmt.Println("Header = ", resp.Header)
	fmt.Println("Body = ", resp.Body)
	fmt.Println("starting reading ...")
	//~ Fetch the reader and start reading ...
	//~ It's the caller's responsibility to close the ReadCloser.
	defer resp.Body.Close() 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Body:", body)
	bodystr := string(body)
	fmt.Printf("len:<%v> : %v\n", len(bodystr), bodystr)
}

func main() {
	//resp, err := http.Post("http://www.01happy.com/demo/accept.php",
	//	"application/x-www-form-urlencoded",
	//	strings.NewReader("name=cjb"),
	//)
	resp, err := http.Post("http://localhost:12345/",
			"application/x-www-form-urlencoded",
			strings.NewReader("onewaystreet"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp)
	//~ OK. We need to analyze http.Response
	fmt.Println()
	printResp(resp)
}
