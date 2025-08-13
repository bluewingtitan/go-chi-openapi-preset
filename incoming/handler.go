package incoming

import (
	"context"
	"github.com/bluewingtitan/go-chi-openapi-preset/core"
)

// ReqHandler todo: rename
type ReqHandler struct {
	Service core.Service
}

func (r ReqHandler) GetExample(_ context.Context, _ GetExampleRequestObject) (GetExampleResponseObject, error) {
	return GetExample2XXJSONResponse{
		Body:       r.Service.GetExample(),
		StatusCode: 0,
	}, nil

}

func NewReqHandler(service core.Service) *ReqHandler {
	return &ReqHandler{
		Service: service,
	}
}
