package failed

// Option ...
type Option func(*Failed)

// WithHTTPStatus ...
func WithHTTPStatus(HTTPStatus int) Option {
	return func(failed *Failed) {
		failed.HTTPStatus = HTTPStatus
	}
}

// WithResponseCode ...
func WithResponseCode(responseCode string) Option {
	return func(failed *Failed) {
		failed.ResponseCode = responseCode
	}
}

// WithResponseDesc ...
func WithResponseDesc(responseDesc string) Option {
	return func(failed *Failed) {
		failed.ResponseDesc = responseDesc
	}
}
