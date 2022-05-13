package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey, to shorten url",
		})
	})
	
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("failed to start the server - Error: %v", err))
	}
}
