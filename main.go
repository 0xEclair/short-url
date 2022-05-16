package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"short-url/api"
	"short-url/cache"
)

func main() {
	cache.InitializeStore()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey, to shorten url",
		})
	})

	r.POST("/api/v1/save", api.ShortUrl)

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("failed to start the server - Error: %v", err))
	}
}
