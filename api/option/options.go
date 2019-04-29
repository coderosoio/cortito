package option

import (
	accountProto "account/proto/account"
	shortenerProto "shortener/proto/shortener"
)

type Options struct {
	UserService    accountProto.UserService
	AuthService    accountProto.AuthService
	LinkService    shortenerProto.LinkService
	AllowedHosts   []string
	AllowedMethods []string
	AllowedHeaders []string
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

func WithAllowedHosts(allowedHosts []string) Option {
	return func(options *Options) {
		options.AllowedHosts = allowedHosts
	}
}

func WithAllowedMethods(allowedMethods []string) Option {
	return func(options *Options) {
		options.AllowedMethods = allowedMethods
	}
}

func WithAllowedHeaders(allowedHeaders []string) Option {
	return func(options *Options) {
		options.AllowedHeaders = allowedHeaders
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
