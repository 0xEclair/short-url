package api

import (
	"github.com/gin-gonic/gin"

	"short-url/service"
)

func ShortUrl(c *gin.Context)  {
	var service service.UrlCreationSerivice
	if err := c.ShouldBind(&service); err == nil {
		service.Save(c)
	} else {
		c.JSON(200, err.Error())
	}
}