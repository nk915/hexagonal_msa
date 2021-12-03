package endpoint

import (
	"context"

	kitendpoint "github.com/go-kit/kit/endpoint"
	imple "local-testing.com/nk915/implementation"
)

//// Endpoints holds all Go kit enpoints for the SaaS Service.
//func Endpoints struct {
//	Create 	kitendpoint.Endpoint
//	GetByID	kitendpoint.Endpoint
//}
//
//// MakeEndpoints initializes all Go kit endpoints for the SaaS Service.
//func MakeEndpoints(svc imple.Service) Endpoints {
//	return Endpoints{
//		Create:		makeCreateEndpoint(svc),
//		GetByID:	makeGetByIDEndpoints(svc),
//	}
//}

// HTTP POST: Create Endpoint
func MakeCreateEndpoint(svc imple.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := svc.Create(ctx, req.SaaS)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

// HTTP GET: GetByID Endpoint
func MakeGetByIDEndpoints(svc imple.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		saasRes, err := svc.GetByID(ctx, req.ID)
		return GetByIDResponse{SaaS: saasRes, Err: err}, nil
	}
}
