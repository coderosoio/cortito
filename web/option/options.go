package option

import (
	"common/keyvalue"
	shortenerProto "shortener/proto/shortener"
)

type Options struct {
	LinkService     shortenerProto.LinkService
	KeyValueStorage keyvalue.Storage
}

type Option func(*Options)

func WithLinkService(linkService shortenerProto.LinkService) Option {
	return func(options *Options) {
		options.LinkService = linkService
	}
}

func WithKeyValueStorage(storage keyvalue.Storage) Option {
	return func(options *Options) {
		options.KeyValueStorage = storage
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
