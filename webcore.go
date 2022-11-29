package shirley

import (
	"net/http"
)

// HandlerFunc defines the request handler
type HandlerFunc func(c *Context)

// Engine implement the interface of ServerHTTP
type Engine struct {
	router *Router
}

// New is the constructor of web.engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// Run defines the method to start a http server
func (engine *Engine) Run(address string) (err error) {
	return http.ListenAndServe(address, engine)
}
