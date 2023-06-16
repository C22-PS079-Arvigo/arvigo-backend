package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yusufwib/arvigo-backend/utils"
)

func ConnectRedis() (*redis.Client, error) {
	var (
		redisHost     = os.Getenv("REDIS_HOST")
		redisPort     = os.Getenv("REDIS_PORT")
		redisUsername = os.Getenv("REDIS_USERNAME")
		redisPassword = os.Getenv("REDIS_PASSWORD")
		redisDB       = os.Getenv("REDIS_DB")
	)

	if redisHost == "" {
		redisHost = "34.101.154.8"
	}

	if redisHost == "" {
		redisPort = "6379"
	}

	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Username: redisUsername,
		Password: redisPassword,
		DB:       utils.StrToInt(redisDB, 0),
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

	// Test Connection And Auth with PING
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping command: %w", err)
	}

	return client, nil
}
