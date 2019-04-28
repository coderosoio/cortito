package option

import (
	shortenerProto "shortener/proto/shortener"
)

type Options struct {
	LinkService shortenerProto.LinkService
}

type Option func(*Options)

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
