package database

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func setup_redis() *redis.Client {
	envs := godotenv.Load(".env")
	if envs != nil {
		log.Fatal("Failed to load environment variables")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_uri"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})
	return client
}
