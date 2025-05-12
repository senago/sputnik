package closer

import "time"

type callback struct {
	config callbackConfig
	fn     func() error
}

func newCallback(fn func() error, opts ...CallbackOption) callback {
	return callback{
		config: newCallbackConfig(opts...),
		fn:     fn,
	}
}

type callbackConfig struct {
	name    string
	timeout time.Duration
}

// CallbackOption опция callback.
type CallbackOption func(*callbackConfig)

func newCallbackConfig(opts ...CallbackOption) callbackConfig {
	config := callbackConfig{
		name:    "closer",
		timeout: 10 * time.Second,
	}

	for i := range opts {
		opts[i](&config)
	}

	return config
}

// WithCallbackName задает name.
//
// По умолчанию: closer.
func WithCallbackName(name string) CallbackOption {
	return func(cc *callbackConfig) { cc.name = name }
}

// WithCallbackTimeout задает timeout.
//
// По умолчанию: 10s.
func WithCallbackTimeout(timeout time.Duration) CallbackOption {
	return func(cc *callbackConfig) { cc.timeout = timeout }
}
