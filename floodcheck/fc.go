package floodcheck

import (
	"context"
	"strconv"
	"task/config"
	"task/db"
	"time"

	"github.com/redis/go-redis/v9"
)

type FC struct {
	client *redis.Client
	N      int64
	K      int64
}

func New(config *config.Config) *FC {
	return &FC{
		client: db.GetReddisClient(config),
		N:      config.Flood.N,
		K:      config.Flood.K,
	}
}

func (obj *FC) Check(ctx context.Context, userID int64) (bool, error) {
	time := time.Now().Unix()
	obj.client.RPush(ctx, strconv.FormatInt(userID, 10), time)
	vals, err := obj.client.LRange(ctx, strconv.FormatInt(userID, 10), 0, -1).Result()
	if err != nil {
		return false, err
	}
	end := int64(len(vals) - 1)
	for idx := int64(len(vals) - 1); idx >= 0; idx-- {
		val, err := strconv.ParseInt(vals[idx], 10, 64)
		if err != nil {
			return false, err
		}
		if val < time-obj.K {
			end = idx + 1
			break
		}
	}
	return !(int64(len(vals))-end > obj.N), nil
}
