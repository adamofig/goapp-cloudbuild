package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(r *gin.Context) {
		html := `

		<body style="background-color: #2D0C62; color: white; text-align: center; padding: 200px;">
			<h1>Cloud Build - Golang </h1>
			<hr>
			<h3> Bienvenido Go APP  </h3>
			<h4>  <a href="/ping"> Do Ping </a> </h4>
			<h4>Y no olvides suscribirte a DATACLOUDER</h4>
		</body>
	`

		r.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
