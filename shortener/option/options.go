package option

import (
	"common/keyvalue"
	"shortener/repository"
)

type Options struct {
	LinkRepository  *repository.LinkRepository
	KeyValueStorage keyvalue.Storage
}

type Option func(*Options)

func WithLinkRepository(linkRepository *repository.LinkRepository) Option {
	return func(options *Options) {
		options.LinkRepository = linkRepository
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
