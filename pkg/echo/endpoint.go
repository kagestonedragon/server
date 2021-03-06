// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
//go:generate easyjson -all endpoint.go
package echo

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	_ "github.com/mailru/easyjson/gen"
)

//easyjson:json
type Echo struct {
	Id       uint32 `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Reminder string `json:"reminder,omitempty"`
}

//easyjson:json
type GetEchoListRequest struct {
}

//easyjson:json
type GetEchoListResponse struct {
	Echos []Echo `json:"echos,omitempty"`
	Err   string `json:"err,omitempty"`
}

//easyjson:skip
type endpoints struct {
	GetEchoEndpoint endpoint.Endpoint
}

func (e endpoints) GetEcho(ctx context.Context, req *GetEchoListRequest) (resp *GetEchoListResponse, err error) {
	response, err := e.GetEchoEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(GetEchoListResponse)
	return &r, err
}

func makeGetEchoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEchoListRequest)
		return s.GetEcho(ctx, &req)
	}
}
