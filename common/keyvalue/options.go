package keyvalue

type Options struct {
	ConnectionName string
	Namespace      string
}

type Option func(*Options)

func WithConnectionName(connectionName string) Option {
	return func(options *Options) {
		options.ConnectionName = connectionName
	}
}

func WithNamespace(namespace string) Option {
	return func(options *Options) {
		options.Namespace = namespace
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
