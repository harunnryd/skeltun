// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package customwriter

import (
	"encoding/json"
	pkgerrors "github.com/harunnryd/skeltun/internal/pkg/errors"
	"github.com/harunnryd/skeltun/internal/pkg/http/transporter"
	"github.com/harunnryd/skeltun/internal/pkg/http/transporter/failed"
	"github.com/harunnryd/skeltun/internal/pkg/http/transporter/success"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
)

// ICustomWriter ...
type ICustomWriter interface {
	WriteSuccess(w http.ResponseWriter, r *http.Request, data interface{})
	WriteError(w http.ResponseWriter, r *http.Request, err error)
}

// CustomWriter ...
type CustomWriter struct{}

// New ...
func New(opts ...Option) ICustomWriter {
	custwriter := new(CustomWriter)
	for _, opt := range opts {
		opt(custwriter)
	}
	return custwriter
}

// WriteSuccess ...
func (customWriter *CustomWriter) WriteSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	response := success.New(
		success.WithHTTPStatus(http.StatusOK),
		success.WithData(data),
		success.WithResponseCode("000000"),
		success.WithResponseDesc("Success"),
	)
	customWriter.write(w, r, response, response.GetHTTPStatus())
}

func (customWriter *CustomWriter) write(w http.ResponseWriter, r *http.Request, response interface{}, httpStatus int) {
	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to unmarshal"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// wiring the meta data in header
	customWriter.wireMetadata(w, r)
	w.WriteHeader(httpStatus)
	_, _ = w.Write(res)
}

func (customWriter *CustomWriter) wireMetadata(w http.ResponseWriter, r *http.Request) {
	if limit := r.URL.Query().Get("limit"); len(strings.TrimSpace(limit)) != 0 {
		w.Header().Set("X-Param-Limit", limit)
	}
	if skip := r.URL.Query().Get("skip"); len(strings.TrimSpace(skip)) != 0 {
		w.Header().Set("X-Param-Skip", skip)
	}
}

// WriteError ...
func (customWriter *CustomWriter) WriteError(w http.ResponseWriter, r *http.Request, err error) {
	response := transporter.ErrUnknown

	if _, ok := err.(*net.OpError); ok {
		response = transporter.ErrDBConnection
	}

	if _, ok := err.(*pq.Error); ok {
		response = transporter.ErrDatabaseSQL
	}

	if _, ok := err.(*pgconn.PgError); ok {
		response = failed.New(
			failed.WithHTTPStatus(http.StatusInternalServerError),
			failed.WithResponseCode("00005"),
			failed.WithResponseDesc(err.Error()),
		)
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		response = transporter.ErrInternal
	}

	if _, ok := err.(*jwt.ValidationError); ok {
		response = transporter.ErrInvalidJWtAuth
	}

	if _, ok := err.(*pkgerrors.ValidationError); ok {
		response = failed.New(
			failed.WithHTTPStatus(http.StatusInternalServerError),
			failed.WithResponseCode("0008"),
			failed.WithResponseDesc(err.Error()),
		)
	}

	if _, ok := err.(*pkgerrors.URLNotFoundError); ok {
		response = transporter.ErrURLNotFound
	}

	if _, ok := err.(*pkgerrors.TimeoutError); ok {
		response = transporter.ErrTimeout
	}

	log.Println("CustomWriter message:", err.Error())
	log.Println("CustomWriter type:", reflect.TypeOf(err))

	customWriter.write(w, r, response, response.GetHTTPStatus())
}
