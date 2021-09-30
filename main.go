package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(r *gin.Context) {
		html := `
	<h3> Bienvenido Go APP  </h3>
	<h4> Do Ping  <a href="/ping"> /ping </a> </h4>
	`

		r.Data(200, "text/html; charset=utf-8", []byte(html))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
