package option

import (
	accountProto "account/proto/account"
	shortenerProto "shortener/proto/shortener"
)

type Options struct {
	UserService accountProto.UserService
	AuthService accountProto.AuthService
	LinkService shortenerProto.LinkService
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

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
