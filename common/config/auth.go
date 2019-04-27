package config

import "time"

const (
	JWTAuthStrategy AuthStrategy = "jwt"
)

type AuthStrategy string

type Auth struct {
	Strategy   AuthStrategy  `default:"jwt"`
	Issuer     string        `default:"coderoso.io"`
	Audience   string        `default:"coderoso.io"`
	Expire     time.Duration `default:"72"`
	PrivateKey string        `yaml:"private_key"`
	PublicKey  string        `yaml:"public_key"`
}
