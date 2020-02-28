package success

// Option ...
type Option func(*Success)

// WithHTTPStatus ...
func WithHTTPStatus(HTTPStatus int) Option {
	return func(success *Success) {
		success.HTTPStatus = HTTPStatus
	}
}

// WithResponseCode ...
func WithResponseCode(responseCode string) Option {
	return func(success *Success) {
		success.ResponseCode = responseCode
	}
}

// WithResponseDesc ...
func WithResponseDesc(responseDesc string) Option {
	return func(success *Success) {
		success.ResponseDesc = responseDesc
	}
}

// WithData ...
func WithData(data interface{}) Option {
	return func(success *Success) {
		success.Data = data
	}
}
