package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("https://www.jamesandnhuanh.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
