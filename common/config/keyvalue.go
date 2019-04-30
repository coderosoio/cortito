package config

import (
	"bytes"
	"fmt"
)

const (
	RedisKeyValueAdapter KeyValueAdapter = "redis"
)

type KeyValueAdapter string

type KeyValue struct {
	Adapter  KeyValueAdapter `default:"redis"`
	Hostname string          `default:"redis"`
	Username string          `default:""`
	Password string          `default:""`
	Port     int             `default:"6379"`
	Params   map[string]interface{}
}

func (k *KeyValue) URL(withAdapter bool) string {
	var buffer bytes.Buffer

	if withAdapter {
		buffer.WriteString(fmt.Sprintf("%s://", k.Adapter))
	}
	if len(k.Username) > 0 {
		buffer.WriteString(k.Username)
	}
	if len(k.Password) > 0 {
		buffer.WriteString(fmt.Sprintf(":%s", k.Password))
	}
	return buffer.String()
}
