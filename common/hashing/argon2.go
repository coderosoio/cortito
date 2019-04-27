package hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argon2HashingStrategy struct {
	params *Params
}

func (s *argon2HashingStrategy) GenerateFromPassword(password string) (string, error) {
	salt, err := generateRandomBytes(s.params.SaltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, s.params.Iterations, s.params.Memory, s.params.Parallelism, s.params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, s.params.Memory, s.params.Iterations, s.params.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (s *argon2HashingStrategy) ComparePasswordAndHash(password string, encodedHash string) (bool, error) {
	params, salt, hash, err := s.decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func (s *argon2HashingStrategy) decodeHash(encodedHash string) (params *Params, salt []byte, hash []byte, err error) {
	values := strings.Split(encodedHash, "$")
	if len(values) != 6 {
		return nil, nil, nil, fmt.Errorf("the encoded hash is not in the correct format")
	}

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, fmt.Errorf("incompatible version of argon2")
	}

	params = &Params{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
