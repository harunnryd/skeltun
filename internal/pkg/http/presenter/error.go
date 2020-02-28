package presenter

import (
	"net/http"
	"skeltun/internal/pkg/http/presenter/failed"
)

// ErrUnknown ...
var ErrUnknown failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusInternalServerError),
	failed.WithResponseCode("00001"),
	failed.WithResponseDesc("Unknown error"),
)

// ErrDBConnection ...
var ErrDBConnection failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusUnauthorized),
	failed.WithResponseCode("00002"),
	failed.WithResponseDesc("Database connection error"),
)

// ErrInvalidHeader ...
var ErrInvalidHeader failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusBadRequest),
	failed.WithResponseCode("00003"),
	failed.WithResponseDesc("Invalid/Incomplete header"),
)

// ErrInvalidHeaderSignature ...
var ErrInvalidHeaderSignature failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusBadRequest),
	failed.WithResponseCode("00004"),
	failed.WithResponseDesc("Invalid header signature"),
)

// ErrDatabaseSQL ...
var ErrDatabaseSQL failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusInternalServerError),
	failed.WithResponseCode("00005"),
	failed.WithResponseDesc("Database sql error"),
)

// ErrURLNotFound ...
var ErrURLNotFound failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusNotFound),
	failed.WithResponseCode("00006"),
	failed.WithResponseDesc("URL not found"),
)
