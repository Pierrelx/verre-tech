package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/PierreLx/verre-tech-ms/store"
	"github.com/PierreLx/verre-tech-ms/store/transport"
)

var (
	//ErrBadRouting correspond a une erreur de routage
	ErrBadRouting = errors.New("Bad routing")
)

//NewService instancie les redirections http
func NewService(svcEndpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger) http.Handler {
	var (
		r            = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)

	r.Methods("POST").Path("/stores").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/stores/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByID,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Store); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	idInt, err := strconv.Atoi(id)
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.GetByIDRequest{ID: idInt}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
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
	case store.ErrStoreNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
