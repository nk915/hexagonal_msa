package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/nk915/k8s_msa_example/services/data"
)

func makeGetSaasInfoEndpoint(_svc data.ISaasService) endpoint.Endpoint {
	return func(_ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(SaasServiceRequest)
		saasInfo, err := _svc.GetSaasInfo(_ctx, req.Id)

		if err != nil {
			return SaasServiceResponse{"", err.Error()}, err
		}
		return SaasServiceResponse{saasInfo, ""}, err
	}
}

type SaasServiceRequest struct {
	Id string `json:"id"`
}

type SaasServiceResponse struct {
	Id    string `json:"id"`
	Email string `json:"id"`
	Tmp   string `json:"id"`
}
