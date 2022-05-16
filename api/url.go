package api

import (
	"github.com/gin-gonic/gin"

	"short-url/service"
)

func Save(c *gin.Context)  {
	var service service.UrlCreationSerivice
	if err := c.ShouldBind(&service); err == nil {
		service.Save(c)
	} else {
		c.JSON(200, err.Error())
	}
}

func Retrieve(c *gin.Context) {
	var service service.ShortUrlService
	if err := c.ShouldBind(&service); err == nil {
		service.OriginalUrl(c)
	} else {
		c.JSON(200, err.Error())
	}
}