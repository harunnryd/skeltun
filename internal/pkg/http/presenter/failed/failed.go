package failed

// IFailed ...
type IFailed interface {
	GetResponseCode() string
	GetResponseDesc() string
	GetHTTPStatus() int
	Error() string
}

// Failed ...
type Failed struct {
	ResponseCode string `json:"response-code"`
	ResponseDesc string `json:"response-desc"`
	HTTPStatus   int    `json:"-"`
}

// New ...
func New(opts ...Option) IFailed {
	failed := new(Failed)
	for _, opt := range opts {
		opt(failed)
	}
	return failed
}

// GetResponseCode ...
func (failed *Failed) GetResponseCode() string {
	return failed.ResponseCode
}

// GetResponseDesc ...
func (failed *Failed) GetResponseDesc() string {
	return failed.ResponseDesc
}

// GetHTTPStatus ...
func (failed *Failed) GetHTTPStatus() int {
	return failed.HTTPStatus
}

// Error ...
func (failed *Failed) Error() string {
	return failed.ResponseDesc
}
