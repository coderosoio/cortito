package hashing

const (
	// MemoryDefault default value for `Params.Memory`.
	MemoryDefault = 64 * 1024
	// IterationsDefault default value for `Params.Iterations`.
	IterationsDefault = 3
	// ParallelismDefault default value for `Params.Parallelism`.
	ParallelismDefault = 2
	// SaltLengthDefault default value for `Params.SaltLength`.
	SaltLengthDefault = 16
	// KeyLengthDefault default value for `Params.KeyLength`.
	KeyLengthDefault = 32

	// BCryptHashingStrategy implements hashing using BCrypt.
	BCryptHashingStrategy Strategy = iota
	// Argon2HashingStrategy implements hashing using Argon2.
	Argon2HashingStrategy
)

// Strategy identifies a hashing strategies.
type Strategy int

// Params contains the parameters needed for hashing passwords.
type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// Hashing is the interface implemented by hashing strategies.
type Hashing interface {
	GenerateFromPassword(password string) (string, error)
	ComparePasswordAndHash(password string, encodedHash string) (bool, error)
}

// NewHashingStrategy returns an `Hashing` implementation.
func NewHashingStrategy(strategy Strategy, params *Params) Hashing {
	if params == nil {
		params = &Params{
			Memory:      MemoryDefault,
			Iterations:  IterationsDefault,
			Parallelism: ParallelismDefault,
			SaltLength:  SaltLengthDefault,
			KeyLength:   KeyLengthDefault,
		}
	}
	switch strategy {
	case BCryptHashingStrategy:
		return &bcryptHashingStrategy{params: params}
	case Argon2HashingStrategy:
		fallthrough
	default:
		return &argon2HashingStrategy{params: params}
	}
}
