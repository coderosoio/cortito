package connection

import (
	"fmt"

	"github.com/go-redis/redis"

	cfg "common/config"
)

func getRedisConnection(settings *cfg.KeyValue) (conn *redis.Client, err error) {
	conn = redis.NewClient(&redis.Options{
		Addr:     settings.URL(false),
		Password: settings.Password,
	})
	if _, err = conn.Ping().Result(); err != nil {
		return
	}
	database := "0"
	if db, found := settings.Params["database"]; found {
		database = fmt.Sprintf("%d", db)
	}
	_, err = conn.Do("SELET", database).Result()
	return
}
