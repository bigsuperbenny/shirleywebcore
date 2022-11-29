package main

import (
	"net/http"

	"github.com/bigsuperbenny/shirley"
)

func main() {

	sh83 := shirley.New()

	sh83.GET("/", func(c *shirley.Context) {
		c.HTML(http.StatusOK, "<p>this is shirley core</p>")
	})

	sh83.GET("/hello", func(c *shirley.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	sh83.POST("/login", func(c *shirley.Context) {
		c.JSON(http.StatusOK, shirley.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	sh83.Run(":9999")
}
