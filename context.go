package shirley

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	//request info
	Path   string
	Method string

	//response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Encapsulation -Get request form data
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Encapsulation -Get request query data
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Setting httpcode
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// setting reponse header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// Write string data to response
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Write json data to response
func (c *Context) JSON(code int, jsonObj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(jsonObj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Write byte data to response
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// Write html data to response
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
