package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9999", nil)

}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%s", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}
