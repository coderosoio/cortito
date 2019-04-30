package keyvalue

import (
	"common/connection"

	"github.com/go-redis/redis"
)

type RedisKeyValueStorage struct {
	connectionName string
	client         *redis.Client
}

func NewRedisKeyValueStorage(options *Options) (Storage, error) {
	storage := &RedisKeyValueStorage{
		connectionName: options.ConnectionName,
	}
	return storage, nil
}

func (s *RedisKeyValueStorage) Connect() (err error) {
	if s.client == nil {
		if s.client, err = connection.GetKeyValueStorageConnection(s.connectionName); err != nil {
			return
		}
	}
	return
}

func (s *RedisKeyValueStorage) Get(key string, defaultValue string) (string, error) {
	value, err := s.client.Get(key).Result()
	if err == redis.Nil {
		value = defaultValue
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func (s *RedisKeyValueStorage) GetIn(key string, field string, defaultValue string) (string, error) {
	value, err := s.client.HGet(key, field).Result()
	if err == redis.Nil {
		value = defaultValue
	} else if err != nil {
		return "", nil
	}
	return value, nil
}

func (s *RedisKeyValueStorage) Set(key string, value string) (err error) {
	err = s.client.Set(key, value, 0).Err()
	return
}

func (s *RedisKeyValueStorage) SetIn(key string, field string, value string) (err error) {
	err = s.client.HSet(key, field, value).Err()
	return
}

func (s *RedisKeyValueStorage) Remove(key string) (err error) {
	err = s.client.Del(key).Err()
	return
}

func (s *RedisKeyValueStorage) RemoveIn(key string, field string) (err error) {
	err = s.client.HDel(key, field).Err()
	return
}
