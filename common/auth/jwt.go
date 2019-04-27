package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"

	commonConfig "common/config"
)

type jwtAuthStrategy struct {
	settings   *commonConfig.Auth
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func newJWTAuthStrategy(settings *commonConfig.Auth) (Auth, error) {
	key, err := ioutil.ReadFile(settings.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error reading JWT private key: %v", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return nil, fmt.Errorf("error parsing JWT private key: %v", err)
	}

	key, err = ioutil.ReadFile(settings.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("error reading JWT public key: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(key)

	strategy := &jwtAuthStrategy{
		settings:   settings,
		publicKey:  publicKey,
		privateKey: privateKey,
	}

	return strategy, nil
}

func (s *jwtAuthStrategy) GenerateToken(data map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	now := time.Now()
	claims := jwt.MapClaims{
		"aud": s.settings.Audience,
		"exp": now.Add(s.settings.Expire * time.Hour).Unix(),
		"iss": s.settings.Issuer,
		"iat": now,
	}
	for key, value := range data {
		claims[key] = value
	}
	token.Claims = claims

	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}

func (s *jwtAuthStrategy) VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return s.publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	return token.Claims.(jwt.MapClaims), nil
}
