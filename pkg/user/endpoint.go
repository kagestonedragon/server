// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
//go:generate easyjson -all endpoint.go
package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	_ "github.com/mailru/easyjson/gen"
)

// AddUserRequest AddUser
//easyjson:json
type AddUserRequest struct {
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
}

// DeleteUserByIdRequest DeleteUser
//easyjson:json
type DeleteUserByIdRequest struct {
	Id uint64 `json:"id,omitempty"`
}

//easyjson:json
type DeleteUserByIdResponse struct {
}

// GetUserByIdRequest GetUser
//easyjson:json
type GetUserByIdRequest struct {
	Id uint64 `json:"id,omitempty"`
}

// GetUserListRequest GetUserList
//easyjson:json
type GetUserListRequest struct {
}

//easyjson:json
type GetUserListResponse []User

//easyjson:json
type UpdateUserResponse struct {
}

//easyjson:json
type User struct {
	Id     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
}

//easyjson:skip
type endpoints struct {
	AddUserEndpoint        endpoint.Endpoint
	GetUserListEndpoint    endpoint.Endpoint
	UpdateUserEndpoint     endpoint.Endpoint
	DeleteUserByIdEndpoint endpoint.Endpoint
	GetUserByIdEndpoint    endpoint.Endpoint
}

func (e endpoints) AddUser(ctx context.Context, req *AddUserRequest) (resp *User, err error) {
	response, err := e.AddUserEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(User)
	return &r, err
}

func (e endpoints) GetUserList(ctx context.Context, req *GetUserListRequest) (resp *GetUserListResponse, err error) {
	response, err := e.GetUserListEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(GetUserListResponse)
	return &r, err
}

func (e endpoints) UpdateUser(ctx context.Context, req *User) (resp *UpdateUserResponse, err error) {
	response, err := e.UpdateUserEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(UpdateUserResponse)
	return &r, err
}

func (e endpoints) DeleteUserById(ctx context.Context, req *DeleteUserByIdRequest) (resp *DeleteUserByIdResponse, err error) {
	response, err := e.DeleteUserByIdEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(DeleteUserByIdResponse)
	return &r, err
}

func (e endpoints) GetUserById(ctx context.Context, req *GetUserByIdRequest) (resp *User, err error) {
	response, err := e.GetUserByIdEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(User)
	return &r, err
}

func makeAddUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)
		return s.AddUser(ctx, &req)
	}
}

func makeGetUserListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserListRequest)
		return s.GetUserList(ctx, &req)
	}
}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		return s.UpdateUser(ctx, &req)
	}
}

func makeDeleteUserByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserByIdRequest)
		return s.DeleteUserById(ctx, &req)
	}
}

func makeGetUserByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIdRequest)
		return s.GetUserById(ctx, &req)
	}
}
