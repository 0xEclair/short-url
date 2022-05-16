package model

import (
	"short-url/cache"
)

func Init() {
	cache.InitializeRedis()
}
