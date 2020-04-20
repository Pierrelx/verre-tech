package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/PierreLx/verre-tech-ms/store"
)

//Endpoints représente un go kit endpoint
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
}

//MakeEndpoints instancie les endpoints
func MakeEndpoints(s store.Service) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(s),
		GetByID: makeGetByIDEndpoint(s),
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
