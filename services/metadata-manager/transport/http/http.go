package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"local-testing.com/nk915/data"
	"local-testing.com/nk915/endpoint"
	"local-testing.com/nk915/implementation"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func NewHttpServer(_svc data.ISaasService, _logger kitlog.Logger) *mux.Router {

	// set-up router and initialize http endpoints
	var (
		r            = mux.NewRouter()
		errorLoger   = kithttp.ServerErrorLogger(_logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)

	options := []kithttp.ServerOption{
		errorLoger,
		errorEncoder,
	}

	// HTTP Get - /services/{id}
	r.Methods("GET").Path("/services/{id}").Handler(kithttp.NewServer(
		endpoint.makeGetSaasInfoEndpoint(_svc),
		decodeGetSaasServiceRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetSaasServiceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoint.SaasServiceRequest{Id: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case implementation.ErrEmpty:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
