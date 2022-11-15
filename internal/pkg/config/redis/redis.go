package redis

import (
	"context"
	rv8 "github.com/go-redis/redis/v8"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"log"
)

var rdb *rv8.Client

func init() {
	dsn := config.GetString("redis.default.host") + ":" + config.GetString("redis.default.port")
	dbNo := config.GetInt("redis.default.database")
	if dbNo < 0 || dbNo > 16 {
		dbNo = 0
	}

	rdb = rv8.NewClient(&rv8.Options{
		Addr: dsn,
		DB:   dbNo,
	})

	if rdb.Ping(context.Background()).Err() != nil {
		log.Fatalf("redis connect err!")
	}
}

func Get() *rv8.Client {
	return rdb
}
