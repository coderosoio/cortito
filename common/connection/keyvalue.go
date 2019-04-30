package connection

import (
	"fmt"

	"github.com/go-redis/redis"

	cfg "common/config"
)

func GetKeyValueStorageConnection(connectionName string) (conn *redis.Client, err error) {
	config, err := cfg.GetConfig()
	if err != nil {
		return nil, err
	}
	settings, found := config.KeyValue[connectionName]
	if !found {
		return nil, fmt.Errorf("could not get connection settings for %s", connectionName)
	}
	switch settings.Adapter {
	case cfg.RedisKeyValueAdapter:
		fallthrough
	default:
		conn, err = getRedisConnection(settings)
	}
	return
}
