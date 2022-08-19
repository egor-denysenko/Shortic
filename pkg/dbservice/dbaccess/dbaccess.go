package dbaccess

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type UrlCacheConnection struct {
	cacheconnection *redis.Client
}

func NewUrlCache() *UrlCacheConnection {
	return &UrlCacheConnection{
		cacheconnection: nil,
	}
}

// Connection to redis database
func (uc *UrlCacheConnection) Connect() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisAddr"), //URL from env tables
		Password: os.Getenv("RedisPass"), // database password
		DB:       0,                      // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	uc.cacheconnection = rdb
	return nil
}

// Function that saves the url for 24 hours into redis database
func (uc *UrlCacheConnection) SaveUrl(fullUrl string, shorterUrl string) error {
	err := uc.cacheconnection.Set(context.Background(), shorterUrl, fullUrl, 86400*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (uc *UrlCacheConnection) CheckUrlCollision(shortUrlKey string) bool {
	_, err := uc.cacheconnection.Get(context.Background(), shortUrlKey).Result()
	if err == redis.Nil {
		return true
	}
	return false
}
func (uc *UrlCacheConnection) FindShortenUrl(shortUrlKey string) string {
	fullUrl, err := uc.cacheconnection.Get(context.Background(), shortUrlKey).Result()
	if err == redis.Nil {
		return ""
	}
	return fullUrl
}
