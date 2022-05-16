package service

import (
	"github.com/gin-gonic/gin"

	"short-url/cache"
	"short-url/shortener"
)

type UrlCreationSerivice struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func (service *UrlCreationSerivice) Save(c *gin.Context) {
	gen := shortener.CommonGenerator{}
	short_url := gen.GenerateShortUrl(service.LongUrl, service.UserId)
	cache.SaveUrlMapping(short_url, service.LongUrl, service.UserId)
	c.JSON(200, short_url)
}
