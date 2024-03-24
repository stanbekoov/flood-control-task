package main

import (
	"context"
	"log"
	"task/config"
	"task/floodcheck"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	log.Println("env file loaded succesfully")
}

func main() {
	config := config.New()
	fc := floodcheck.New(config)
	ctx := context.Background()
	fc.Check(ctx, 0)
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
