package presenter

// Meta defines meta format for api format
type Meta struct {
	Version string `json:"version"`
	Status  string `json:"api_status"`
	APIEnv  string `json:"api_env"`
}

// Response ...
type Response struct {
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
	Meta         Meta   `json:"meta"`
}

// SuccessResponse ...
type SuccessResponse struct {
	Response
	Next       *string     `json:"next,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	HTTPStatus int         `json:"-"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Response
	HTTPStatus int `json:"-"`
}

func (e *ErrorResponse) Error() string {
	return e.ResponseDesc
}
