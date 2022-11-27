package shirley

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine implement the interface of ServerHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of web.engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// web engine add route
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	log.Printf("Request %4s - %s", req.Method, req.URL.Path)
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL.Path)
	}
}

// Run defines the method to start a http server
func (engine *Engine) Run(address string) (err error) {
	return http.ListenAndServe(address, engine)
}
