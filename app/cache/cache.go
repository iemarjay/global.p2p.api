package cache

import (
	"global.p2p.api/app"
	"time"
)
import "github.com/go-redis/redis"

type (
	Cache struct {
		redisClient *redis.Client
	}
)

func (c *Cache) Set(key string, value string, expiry time.Duration) error {
	err := c.redisClient.Set(key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	value, err := c.redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cache) GetOrDefault(key string, fail string) (string) {
	value, err := c.redisClient.Get(key).Result()
	if err != nil {
		return fail
	}

	return value
}

func New(env *app.Env) *Cache {
	return &Cache{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     env.Get("REDIS_URL"),
			Password: env.Get("REDIS_PASSWORD"),
		}),
	}
}