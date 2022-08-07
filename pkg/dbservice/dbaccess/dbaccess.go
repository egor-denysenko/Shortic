package dbaccess

import (
	//"os"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type UrlCacheConnection struct{
  cacheconnection *redis.Client
}

func NewUrlCache() *UrlCacheConnection{
  return &UrlCacheConnection{
      cacheconnection: nil,
  }
}

// Connection to redis database
func (uc *UrlCacheConnection) Connect() error{
  rdb := redis.NewClient(&redis.Options{
    //Addr:     os.Getenv("RedisAddr"), //URL from env tables
		//Password: "", // database password
    Addr:    "eu1-desired-loon-38077.upstash.io:38077",
		Password: "326116961f9746f0b7b57b387065f923", // database password
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	uc.cacheconnection = rdb
	return nil
}

// Function that saves the url for 24 hours into redis database
func (uc *UrlCacheConnection) SaveUrl(fullUrl string, shorterUrl string) error{
  err := uc.cacheconnection.Set(nil, fullUrl, shorterUrl, 86400*time.Second).Err()
	if err != nil {
		return err
	}
  return nil
}
