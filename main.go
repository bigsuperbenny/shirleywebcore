package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := Engine{}
	//Function http.ListenAndServe Param
	//Param1:listen address,9999 is address port
	//Param2:Implementation of Web Framework Entry Based on Net/http Standard Library
	log.Fatal(http.ListenAndServe(":9999", &engine))

}

type Engine struct{}

func (*Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%s", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
