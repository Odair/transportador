package transportador

import (
	"context"

	"github.com/go-kit/kit/endpoint"

)

type Endpoints struct {
	CriarEntrega endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CriarEntrega: makeCriarEntregaEndpoint(s),
	}
}

func makeCriarEntregaEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CriarEntregaRequest)
		voucher, err := s.CriarEntrega(ctx, req.Entrega)
		return CriarEntregaResponse{voucher}, err
	}
}
