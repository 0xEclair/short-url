package service

import (
	"github.com/gin-gonic/gin"

	"short-url/cache"
)

type ShortUrlService struct {
	ShortUrl string `json:"short_url" binding:"required"`
}

func (service *ShortUrlService) OriginalUrl(c *gin.Context)  {
	res := cache.RetrieveInitialUrl(service.ShortUrl)
	c.JSON(200, res)
}