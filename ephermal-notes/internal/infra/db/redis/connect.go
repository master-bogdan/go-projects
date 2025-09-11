package redis

import (
	"context"

	"github.com/master-bogdan/ephermal-notes/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Connect(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&cfg.Db.Redis)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
