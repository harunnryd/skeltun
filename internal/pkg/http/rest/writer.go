package rest

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"skeltun/internal/pkg/http/presenter"

	"github.com/go-chi/render"
	"github.com/lib/pq"
)

func respondwithSuccess(w http.ResponseWriter, r *http.Request, psenter interface{}) {
	httpStatus := lookUpSuccess(psenter)
	writeResponse(w, r, psenter, httpStatus)
}

func respondwithError(w http.ResponseWriter, r *http.Request, err error) {
	httpStatus, psenter := lookUpErr(err)
	writeResponse(w, r, psenter, httpStatus)
}

func lookUpErr(err error) (httpStatus int, psenter error) {
	switch err {
	case err.(*net.OpError):
		{
			psenter = presenter.ErrDatabase
			httpStatus = presenter.ErrDatabase.HTTPStatus
		}
	case err.(*pq.Error):
		{
			psenter = presenter.ErrDatabaseAuthFailed
			httpStatus = presenter.ErrDatabaseAuthFailed.HTTPStatus
		}
	default:
		{
			fmt.Printf("[Lookup] Type : %s\n", reflect.TypeOf(err))
			fmt.Printf("[Lookup] Error: %s\n", err.Error())
			psenter = presenter.ErrUnknown
			httpStatus = presenter.ErrUnknown.HTTPStatus
		}
	}
	return
}

func lookUpSuccess(v interface{}) (httpStatus int) {
	if _, ok := v.(presenter.SuccessResponse); ok {
		httpStatus = v.(presenter.SuccessResponse).HTTPStatus
	}
	httpStatus = http.StatusOK
	return
}

func writeResponse(w http.ResponseWriter, r *http.Request, response interface{}, httpStatus int) {
	render.Status(r, httpStatus)
	render.JSON(w, r, response)
}
