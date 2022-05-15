package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type pair struct {
	initial string
	short   string
}

const (
	UserId = "user-id-test-example"
)

var (
	url_1 = &pair{
		initial: "https://www.bilibili.com",
		short:   "NpTbeNTw",
	}
	url_2 = pair{
		initial: "https://www.github.com",
		short:   "GVqTgfdf",
	}
)

func TestCommonGenerator(t *testing.T) {
	generator := CommonGenerator{}
	
	shortUrl_1 := generator.GenerateShortUrl(url_1.initial, UserId)
	
	shortUrl_2 := generator.GenerateShortUrl(url_2.initial, UserId)
	
	assert.Equal(t, shortUrl_1, url_1.short)
	assert.Equal(t, shortUrl_2, url_2.short)
}
