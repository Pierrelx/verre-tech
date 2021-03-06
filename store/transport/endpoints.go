package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/Pierrelx/verre-tech/store"
)

//Endpoints représente un go kit endpoint
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
	Update  endpoint.Endpoint
	Delete  endpoint.Endpoint
	GetAll  endpoint.Endpoint
}

//MakeEndpoints instancie les endpoints
func MakeEndpoints(s store.Service) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(s),
		GetByID: makeGetByIDEndpoint(s),
		Update:  makeUpdateEndpoints(s),
		Delete:  makeDeleteEndpoints(s),
		GetAll:  makeGetAllEndpoints(s),
	}
}

func makeCreateEndpoint(s store.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(ctx, req.Store)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeGetByIDEndpoint(s store.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		storeRes, err := s.GetByID(ctx, req.ID)
		return GetByIDResponse{Store: storeRes, Err: err}, nil
	}
}

func makeUpdateEndpoints(s store.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		storeRes, err := s.UpdateStore(ctx, req.Store)
		return GetByIDResponse{Store: storeRes, Err: err}, nil
	}
}

func makeDeleteEndpoints(s store.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.DeleteStore(ctx, req.ID)
		return DeleteResponse{Err: err}, nil
	}
}

func makeGetAllEndpoints(s store.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		stores, err := s.GetAll(ctx)
		return ListStoreResponse{Stores: stores, Err: err}, nil
	}
}
