package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {

	res, _ := http.Get("http://www.globo.com/")
	fmt.Println("%s", res)
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s",body)
}
