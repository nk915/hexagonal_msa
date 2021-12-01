package handling

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type makeGetSaasInfoEndpoint(_svc saasService) endpoint.Endpoint {
	return func(_ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(saasServiceRequest)
		saasInfo, err := _svc.getSaasInfo(ctx, req.Id)

		if err != nil {
			return 
		}
	}
}



type saasServiceRequest struct {
	Id	string `json:"id"`
}


