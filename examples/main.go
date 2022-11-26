package main

import (
	"fmt"
	"net/http"

	"github.com/BigSuperBenny/ShirleyWebCore"
)

func main() {

	rCore := ShirleyWebCore.New()

	rCore.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	rCore.POST("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	rCore.Run(":9999")
}
