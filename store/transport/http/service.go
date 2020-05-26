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

	"github.com/PierreLx/verre-tech/store"
	"github.com/PierreLx/verre-tech/store/transport"
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

	r.Methods("POST").Path("/store").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/store/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByID,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/stores").Handler(kithttp.NewServer(
		svcEndpoints.GetAll,
		decodeGetAllRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/store/update").Handler(kithttp.NewServer(
		svcEndpoints.Update,
		decodeUpdateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/store/{id}").Handler(kithttp.NewServer(
		svcEndpoints.Delete,
		decodeDeleteRequest,
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

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var store store.Store
	err = decoder.Decode(&store)
	print(r.Body)
	if err != nil {
		return store, err
	}
	println("Magasin : " + store.Name)
	return transport.UpdateRequest{Store: store}, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	idInt, err := strconv.Atoi(id)
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.DeleteRequest{ID: idInt}, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return transport.GetAllRequest{}, nil
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

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
