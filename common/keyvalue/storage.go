package keyvalue

import (
	cfg "common/config"
	"fmt"
)

type Storage interface {
	Set(key string, value string) error
	SetIn(key string, field string, value string) error
	Get(key string, defaultValue string) (string, error)
	GetIn(key string, field string, defaultValue string) (string, error)
	Remove(key string) error
	RemoveIn(key string, field string) error
	Connect() error
}

func NewKeyValueStorage(connectionName string) (Storage, error) {
	config, err := cfg.GetConfig()
	if err != nil {
		return nil, err
	}
	settings, found := config.KeyValue[connectionName]
	if !found {
		return nil, fmt.Errorf("no keyvalue connection found for %s", connectionName)
	}
	options := NewOptions(
		WithConnectionName(connectionName),
	)
	switch settings.Adapter {
	case cfg.RedisKeyValueAdapter:
		fallthrough
	default:
		return NewRedisKeyValueStorage(options)
	}
}
