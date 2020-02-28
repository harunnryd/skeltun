package rest

// Option ...
type Option func(*Rest)

// WithHTTPMethod ...
func WithHTTPMethod(httpMethod string) Option {
	return func(rest *Rest) {
		rest.HTTPMethod = httpMethod
	}
}

// WithPattern ...
func WithPattern(pattern string) Option {
	return func(rest *Rest) {
		rest.Pattern = pattern
	}
}

// WithHandler ...
func WithHandler(h Handler) Option {
	return func(rest *Rest) {
		rest.H = h
	}
}
