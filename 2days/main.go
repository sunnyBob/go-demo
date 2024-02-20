package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	g := gee.New()
	g.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Go!</h1>")
	})
	g.GET("/home", func(c *gee.Context) {
		c.HTML(http.StatusOK, fmt.Sprintf("<h1>Hello %s!</h1>", c.Query("name")))
	})
	g.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"useName": c.PostForm("userName"),
			"passwd":  c.PostForm("passwd"),
		})
	})
	g.Run(":8081")
}
