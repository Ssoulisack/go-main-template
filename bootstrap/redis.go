package bootstrap

import (
	"context"
	"fmt"
	"kkl-v2/core/logs"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitializeRedis(env *Env) *redis.Client {
	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.Redis.Host, env.Redis.Port),
		Password: env.Redis.Password,
		DB:       env.Redis.DB,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("can not connect to redis")
	}

	logs.Info("redis connection success")
	return rdb
}
