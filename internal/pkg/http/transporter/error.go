// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transporter

import (
	"github.com/harunnryd/skeltun/internal/pkg/http/transporter/failed"
	"net/http"
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

// ErrDatabaseSQL ...
var ErrDatabaseSQL failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusInternalServerError),
	failed.WithResponseCode("00005"),
	failed.WithResponseDesc("Database sql error."),
)

// ErrURLNotFound ...
var ErrURLNotFound failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusNotFound),
	failed.WithResponseCode("00006"),
	failed.WithResponseDesc("URL not found."),
)

// ErrInvalidJWtAuth ...
var ErrInvalidJWtAuth failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusBadRequest),
	failed.WithResponseCode("00006"),
	failed.WithResponseDesc("Invalid token/header signature."),
)

// ErrInternal ...
var ErrInternal failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusInternalServerError),
	failed.WithResponseCode("00007"),
	failed.WithResponseDesc("Internal error."),
)

// ErrTimeout ...
var ErrTimeout failed.IFailed = failed.New(
	failed.WithHTTPStatus(http.StatusGatewayTimeout),
	failed.WithResponseCode("00013"),
	failed.WithResponseDesc("Timeout/Cancelation.."),
)
