package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(http.StatusOK)
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", gin.H{})
	})
	r.Run()
	fmt.Println("Hello")
}
