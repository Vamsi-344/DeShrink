package main

import (
	"github.com/Vamsi-344/DeShrink/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/shorten", handlers.ShortURLGenerator)
	r.GET("/:after", handlers.Redirect)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
