package dbaccess

import (
  "os"
  "context"
  "fmt"
	"github.com/go-redis/redis/v8"
)

type UrlCacheConnection struct{
  cacheconection *redis.Client
}

func NewUrlCache() *UrlCacheConnection{
  return &UrlCacheConnection{
      cacheconection: nil,
  }
}

func (uc *UrlCacheConnection) ConnectToCache() error{
  rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisAddr"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	fmt.Println(err)
	if err != nil {
		return err
	}
	uc.cacheconection = rdb
	return nil
}

// Not implemented
func (uc *UrlCacheConnection) SaveUrl() error{
  return nil
}
