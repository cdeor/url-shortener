package database

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func CreateClient(dbNum int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNum,
	})

	return rdb

}
