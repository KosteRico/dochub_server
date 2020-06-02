package redisDb

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

var termsClient *redis.Client
var metaClient *redis.Client

func Init() ([]string, error) {
	pong1, err := initTermsClient()

	if err != nil {
		return nil, err
	}

	pong2, err := initMetaClient()

	if err != nil {
		return nil, err
	}

	return []string{pong1, pong2}, nil
}

func initTermsClient() (string, error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("redis_host"),
		os.Getenv("redis_port"))

	termsClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return termsClient.Ping().Result()
}

func initMetaClient() (string, error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("redis_host"),
		os.Getenv("redis_port"))

	metaClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       1,
	})

	return termsClient.Ping().Result()
}
