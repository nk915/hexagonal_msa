package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	ep "local-testing.com/nk915/endpoint"
	imple "local-testing.com/nk915/implementation"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func NewHttpServer(svc imple.Service, logger kitlog.Logger) *mux.Router {

	// set-up router and initialize http endpoints
	var (
		r            = mux.NewRouter()
		errorLoger   = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)

	options := []kithttp.ServerOption{
		errorLoger,
		errorEncoder,
	}

	// HTTP GET - /services/{id}
	r.Methods("GET").Path("/services/{id}").Handler(kithttp.NewServer(
		ep.MakeGetByIDEndpoints(svc),
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	// HTTP POST - /services
	r.Methods("POST").Path("/services").Handler(kithttp.NewServer(
		ep.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return nil, ErrBadRouting
	}
	return ep.GetByIDRequest{ID: id}, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req ep.CreateRequest
	fmt.Println("TEST TEST", r.Body)
	if e := json.NewDecoder(r.Body).Decode(&req.SaaS); e != nil {
		return nil, e
	}
	fmt.Println("TEST TEST", req)
	return req, nil
}

// TODO: decodeUpdateRequest
//func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
//}

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
	case imple.ErrEmpty:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
