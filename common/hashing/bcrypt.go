package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptHashingStrategy struct {
	params *Params
}

func (s *bcryptHashingStrategy) GenerateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *bcryptHashingStrategy) ComparePasswordAndHash(password string, encodedHash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encodedHash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}
