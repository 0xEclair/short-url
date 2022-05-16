package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeRedis()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.bilibili.com/"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "D3vgAs5oK"

	SaveUrlMapping(shortURL, initialLink, userUUID)
	retrievedURL := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedURL)
}
