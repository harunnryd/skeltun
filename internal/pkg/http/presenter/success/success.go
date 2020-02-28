package success

// ISuccess ...
type ISuccess interface {
	GetResponseCode() string
	GetResponseDesc() string
	GetData() interface{}
	GetHTTPStatus() int
}

// Success ...
type Success struct {
	ResponseCode string      `json:"response-code"`
	ResponseDesc string      `json:"response-desc"`
	Data         interface{} `json:"data,omitempty"`
	HTTPStatus   int         `json:"-"`
}

// New ...
func New(opts ...Option) ISuccess {
	success := new(Success)
	for _, opt := range opts {
		opt(success)
	}
	return success
}

// GetResponseCode ...
func (success *Success) GetResponseCode() string {
	return success.ResponseCode
}

// GetResponseDesc ...
func (success *Success) GetResponseDesc() string {
	return success.ResponseDesc
}

// GetData ...
func (success *Success) GetData() interface{} {
	return success.Data
}

// GetHTTPStatus ...
func (success *Success) GetHTTPStatus() int {
	return success.HTTPStatus
}
