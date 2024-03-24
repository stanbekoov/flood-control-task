package config

import (
	"log"
	"os"
	"strconv"
)

type ReddisConfig struct {
	Address  string
	Password string
	DBName   int
}

type FloodRestrict struct {
	N int64
	K int64
}

type Config struct {
	Red   ReddisConfig
	Flood FloodRestrict
}

func New() *Config {
	return &Config{
		Red: ReddisConfig{
			Address:  getEnv("ADDRESS", ""),
			Password: getEnv("REDDIS_PASSWORD", ""),
			DBName:   int(getEnvInt("DB_NAME", 0)),
		},
		Flood: FloodRestrict{
			N: getEnvInt("N", 0),
			K: getEnvInt("K", 0),
		},
	}
}

func getEnvInt(key string, defaultValue int64) int64 {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	res, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func getEnv(key string, defaultValue string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return val
}
