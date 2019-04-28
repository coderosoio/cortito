package option

import (
	accountProto "account/proto/account"
	"common/config"
	shortenerProto "shortener/proto/shortener"
)

type Options struct {
	UserService   accountProto.UserService
	AuthService   accountProto.AuthService
	LinkService   shortenerProto.LinkService
	SessionSecret string
	SessionStore  config.SessionStore
	SessionName   string
}

type Option func(*Options)

func WithUserService(userService accountProto.UserService) Option {
	return func(options *Options) {
		options.UserService = userService
	}
}

func WithAuthService(authService accountProto.AuthService) Option {
	return func(options *Options) {
		options.AuthService = authService
	}
}

func WithLinkService(linkService shortenerProto.LinkService) Option {
	return func(options *Options) {
		options.LinkService = linkService
	}
}

func WithSessionSecret(sessionSecret string) Option {
	return func(options *Options) {
		options.SessionSecret = sessionSecret
	}
}

func WithSessionStore(sessionStore config.SessionStore) Option {
	return func(options *Options) {
		options.SessionStore = sessionStore
	}
}

func WithSessionName(sessionName string) Option {
	return func(options *Options) {
		options.SessionName = sessionName
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
