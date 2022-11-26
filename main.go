package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/hello", helloHandler)
	//Function http.ListenAndServe Param
	//Param1:listen address,9999 is address port
	//Param2:Implementation of Web Framework Entry Based on Net/http Standard Library
	log.Fatal(http.ListenAndServe(":9999", nil))

}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%s", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}
