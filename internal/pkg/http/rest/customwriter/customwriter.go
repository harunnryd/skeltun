package customwriter

import (
	"encoding/json"
	"net"
	"net/http"
	"skeltun/internal/pkg/http/presenter"
	"skeltun/internal/pkg/http/presenter/success"
	"strings"

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
	cuswriter := new(CustomWriter)
	for _, opt := range opts {
		opt(cuswriter)
	}
	return cuswriter
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
		w.Write([]byte("Failed to unmarshal"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// wiring the meta data in header
	customWriter.wireMetadata(w, r)
	w.WriteHeader(httpStatus)
	w.Write(res)
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
	response := presenter.ErrUnknown

	if _, ok := err.(*net.OpError); ok {
		response = presenter.ErrDBConnection
	}

	if _, ok := err.(*pq.Error); ok {
		response = presenter.ErrDatabaseSQL
	}

	customWriter.write(w, r, response, response.GetHTTPStatus())
}
