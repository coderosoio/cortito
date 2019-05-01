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

func (s *KeyValue) URL(withAdapter bool) string {
	var buffer bytes.Buffer

	if withAdapter {
		buffer.WriteString(fmt.Sprintf("%s://", s.Adapter))
	}
	buffer.WriteString(fmt.Sprintf("%s:%d", s.Hostname, s.Port))
	return buffer.String()
}
