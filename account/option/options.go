package option

import (
	"common/auth"
	"common/hashing"

	"account/repository"
)

// Options is a list of options.
type Options struct {
	UserRepository      *repository.UserRepository
	UserTokenRepository *repository.UserTokenRepository
	HashingStrategy     hashing.Hashing
	AuthStrategy        auth.Auth
}

// Option is a function for setting options.
type Option func(*Options)

// WithUserRepository sets the user repository.
func WithUserRepository(userRepository *repository.UserRepository) Option {
	return func(options *Options) {
		options.UserRepository = userRepository
	}
}

// WithUserTokenRepository sets the user token repository.
func WithUserTokenRepository(userTokenRepository *repository.UserTokenRepository) Option {
	return func(options *Options) {
		options.UserTokenRepository = userTokenRepository
	}
}

// WithHashingStrategy sets the password hashing strategy to use.
func WithHashingStrategy(hashingStrategy hashing.Hashing) Option {
	return func(options *Options) {
		options.HashingStrategy = hashingStrategy
	}
}

// WithAuthStrategy sets the auth strategy to use.
func WithAuthStrategy(authStrategy auth.Auth) Option {
	return func(options *Options) {
		options.AuthStrategy = authStrategy
	}
}

// NewOptions returns a list of default options.
func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
