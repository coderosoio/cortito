package auth

import (
	commonConfig "common/config"
)

type Auth interface {
	GenerateToken(data map[string]interface{}) (token string, err error)
	VerifyToken(token string) (data map[string]interface{}, err error)
}

func NewAuthStrategy() (Auth, error) {
	config, err := commonConfig.GetConfig()
	if err != nil {
		return nil, err
	}
	settings := config.Auth
	switch settings.Strategy {
	case commonConfig.JWTAuthStrategy:
		fallthrough
	default:
		return newJWTAuthStrategy(settings)
	}
}
