package option

import "shortener/repository"

type Options struct {
	LinkRepository *repository.LinkRepository
}

type Option func(*Options)

func WithLinkRepository(linkRepository *repository.LinkRepository) Option {
	return func(options *Options) {
		options.LinkRepository = linkRepository
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
