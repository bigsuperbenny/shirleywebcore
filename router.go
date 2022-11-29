package shirley

import (
	"fmt"
	"log"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

// web engine add route
func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	log.Printf("Request %4s - %s", c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND:%s\n", c.Path)
	}
}
