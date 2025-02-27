package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type CacheSuite struct {
	suite.Suite
	cacheManager Manager
}

func (suite *CacheSuite) SetupTest() {
	suite.cacheManager = NewManager(1 * time.Second)
}

func (suite *CacheSuite) TestSetAndGet() {
	const (
		key        = "testKey"
		data       = "testData"
		expiration = 2 * time.Second
	)

	suite.cacheManager.Set(key, data, expiration)

	cachedData, found := suite.cacheManager.Get(key)

	suite.True(found)
	suite.Equal(data, cachedData)
}

func (suite *CacheSuite) TestExpiredEntryNotReturned() {
	const (
		key        = "testKey"
		data       = "testData"
		expiration = 1 * time.Second
	)

	suite.cacheManager.Set(key, data, expiration)

	// Wait for the entry to expire
	time.Sleep(2 * time.Second)

	cachedData, found := suite.cacheManager.Get(key)

	suite.False(found)
	suite.Nil(cachedData)
}

func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}
