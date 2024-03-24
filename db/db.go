package db

import (
	"task/config"

	"github.com/redis/go-redis/v9"
)

func GetReddisClient(config *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Red.Address,
		Password: config.Red.Password,
		DB:       config.Red.DBName,
	})
	return client
}
