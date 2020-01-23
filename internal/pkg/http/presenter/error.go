package presenter

import "net/http"

// ErrUnknown ...
var ErrUnknown *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00001",
		ResponseDesc: "Unknown error",
	},
	HTTPStatus: http.StatusInternalServerError,
}

// ErrUnauthorized ...
var ErrUnauthorized *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00002",
		ResponseDesc: "You're not authorized",
	},
	HTTPStatus: http.StatusUnauthorized,
}

// ErrInvalidHeader ...
var ErrInvalidHeader *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00003",
		ResponseDesc: "Invalid/Incomplete header",
	},
	HTTPStatus: http.StatusBadRequest,
}

// ErrInvalidHeaderSignature ...
var ErrInvalidHeaderSignature *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00004",
		ResponseDesc: "Invalid header signature",
	},
	HTTPStatus: http.StatusBadRequest,
}

// ErrDatabase ...
var ErrDatabase *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00005",
		ResponseDesc: "Database error",
	},
	HTTPStatus: http.StatusInternalServerError,
}

// ErrURLNotFound ...
var ErrURLNotFound *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00006",
		ResponseDesc: "URL not found",
	},
	HTTPStatus: http.StatusNotFound,
}

// ErrDatabaseAuthFailed ...
var ErrDatabaseAuthFailed *ErrorResponse = &ErrorResponse{
	Response: Response{
		ResponseCode: "00007",
		ResponseDesc: "Authentication failed",
	},
	HTTPStatus: http.StatusInternalServerError,
}
