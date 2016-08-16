

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	response, _ := http.Get("http://www.baidu.com")
	defer response.Body.Close() //~ Read to the end.
	body, _ := ioutil.ReadAll(response.Body) //ReadCloser
	fmt.Println(string(body))
}
