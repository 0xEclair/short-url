package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"short-url/api"
	"short-url/model"
)

func main() {
	model.Init()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey, to shorten url",
		})
	})

	r.POST("/api/v1/save", api.Save)

	r.GET("/api/v1/retrieve", api.Retrieve)

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("failed to start the server - Error: %v", err))
	}
}
