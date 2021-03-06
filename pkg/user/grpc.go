// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package user

import (
	"context"
	"errors"

	pb "github.com/kagestonedragon/server/internal/echopb"
	"github.com/kagestonedragon/server/tools/logging"
	"github.com/kagestonedragon/server/tools/tracing"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	addUser        grpctransport.Handler
	getUserList    grpctransport.Handler
	updateUser     grpctransport.Handler
	deleteUserById grpctransport.Handler
	getUserById    grpctransport.Handler
}

type ContextGRPCKey struct{}

type GRPCInfo struct{}

// NewGRPCServer makes a set of endpoints available as a gRPC userServer.
func NewGRPCServer(ctx context.Context, s Service) pb.UserServiceServer {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "grpc handler", "user")
	tracer := tracing.FromContext(ctx)

	options := []grpctransport.ServerOption{
		// grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(grpcToContext()),
		grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "grpc server", logger)),
		grpctransport.ServerFinalizer(closeGRPCTracer()),
	}

	return &grpcServer{
		addUser: grpctransport.NewServer(
			makeAddUserEndpoint(s),
			decodeGRPCAddUserRequest,
			encodeGRPCAddUserResponse,
			options...,
		),
		getUserList: grpctransport.NewServer(
			makeGetUserListEndpoint(s),
			decodeGRPCGetUserListRequest,
			encodeGRPCGetUserListResponse,
			options...,
		),
		updateUser: grpctransport.NewServer(
			makeUpdateUserEndpoint(s),
			decodeGRPCUpdateUserRequest,
			encodeGRPCUpdateUserResponse,
			options...,
		),
		deleteUserById: grpctransport.NewServer(
			makeDeleteUserByIdEndpoint(s),
			decodeGRPCDeleteUserByIdRequest,
			encodeGRPCDeleteUserByIdResponse,
			options...,
		),
		getUserById: grpctransport.NewServer(
			makeGetUserByIdEndpoint(s),
			decodeGRPCGetUserByIdRequest,
			encodeGRPCGetUserByIdResponse,
			options...,
		),
	}
}

func JoinGRPC(ctx context.Context, s Service) func(*googlegrpc.Server) {
	return func(g *googlegrpc.Server) {
		pb.RegisterUserServiceServer(g, NewGRPCServer(ctx, s))
	}
}

func grpcToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return context.WithValue(ctx, ContextGRPCKey{}, GRPCInfo{})
	}
}

func closeGRPCTracer() grpc.ServerFinalizerFunc {
	return func(ctx context.Context, err error) {
		span := stdopentracing.SpanFromContext(ctx)
		span.Finish()
	}
}

func (s *grpcServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.User, error) {
	_, rep, err := s.addUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.User), nil
}

func (s *grpcServer) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	_, rep, err := s.getUserList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserListResponse), nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, req *pb.User) (*pb.UpdateUserResponse, error) {
	_, rep, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserResponse), nil
}

func (s *grpcServer) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	_, rep, err := s.deleteUserById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserByIdResponse), nil
}

func (s *grpcServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	_, rep, err := s.getUserById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.User), nil
}

func decodeGRPCAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.AddUserRequest)
	if !ok {
		return nil, errors.New("decodeGRPCAddUserRequest wrong request")
	}

	req := PBToAddUserRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func decodeGRPCGetUserListRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.GetUserListRequest)
	if !ok {
		return nil, errors.New("decodeGRPCGetUserListRequest wrong request")
	}

	req := PBToGetUserListRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func decodeGRPCUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.User)
	if !ok {
		return nil, errors.New("decodeGRPCUpdateUserRequest wrong request")
	}

	req := PBToUser(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func decodeGRPCDeleteUserByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.DeleteUserByIdRequest)
	if !ok {
		return nil, errors.New("decodeGRPCDeleteUserByIdRequest wrong request")
	}

	req := PBToDeleteUserByIdRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func decodeGRPCGetUserByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.GetUserByIdRequest)
	if !ok {
		return nil, errors.New("decodeGRPCGetUserByIdRequest wrong request")
	}

	req := PBToGetUserByIdRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func encodeGRPCAddUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*User)
	if !ok {
		return nil, errors.New("encodeGRPCAddUserResponse wrong response")
	}

	return UserToPB(inResp), nil
}

func encodeGRPCGetUserListResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*GetUserListResponse)
	if !ok {
		return nil, errors.New("encodeGRPCGetUserListResponse wrong response")
	}

	return GetUserListResponseToPB(inResp), nil
}

func encodeGRPCUpdateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*UpdateUserResponse)
	if !ok {
		return nil, errors.New("encodeGRPCUpdateUserResponse wrong response")
	}

	return UpdateUserResponseToPB(inResp), nil
}

func encodeGRPCDeleteUserByIdResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*DeleteUserByIdResponse)
	if !ok {
		return nil, errors.New("encodeGRPCDeleteUserByIdResponse wrong response")
	}

	return DeleteUserByIdResponseToPB(inResp), nil
}

func encodeGRPCGetUserByIdResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*User)
	if !ok {
		return nil, errors.New("encodeGRPCGetUserByIdResponse wrong response")
	}

	return UserToPB(inResp), nil
}

func AddUserRequestToPB(d *AddUserRequest) *pb.AddUserRequest {
	if d == nil {
		return nil
	}

	resp := pb.AddUserRequest{
		Name:   d.Name,
		Active: d.Active,
	}

	return &resp
}

func PBToAddUserRequest(d *pb.AddUserRequest) *AddUserRequest {
	if d == nil {
		return nil
	}

	resp := AddUserRequest{
		Name:   d.Name,
		Active: d.Active,
	}

	return &resp
}

func DeleteUserByIdRequestToPB(d *DeleteUserByIdRequest) *pb.DeleteUserByIdRequest {
	if d == nil {
		return nil
	}

	resp := pb.DeleteUserByIdRequest{
		Id: d.Id,
	}

	return &resp
}

func PBToDeleteUserByIdRequest(d *pb.DeleteUserByIdRequest) *DeleteUserByIdRequest {
	if d == nil {
		return nil
	}

	resp := DeleteUserByIdRequest{
		Id: d.Id,
	}

	return &resp
}

func DeleteUserByIdResponseToPB(d *DeleteUserByIdResponse) *pb.DeleteUserByIdResponse {
	if d == nil {
		return nil
	}

	resp := pb.DeleteUserByIdResponse{}

	return &resp
}

func PBToDeleteUserByIdResponse(d *pb.DeleteUserByIdResponse) *DeleteUserByIdResponse {
	if d == nil {
		return nil
	}

	resp := DeleteUserByIdResponse{}

	return &resp
}

func GetUserByIdRequestToPB(d *GetUserByIdRequest) *pb.GetUserByIdRequest {
	if d == nil {
		return nil
	}

	resp := pb.GetUserByIdRequest{
		Id: d.Id,
	}

	return &resp
}

func PBToGetUserByIdRequest(d *pb.GetUserByIdRequest) *GetUserByIdRequest {
	if d == nil {
		return nil
	}

	resp := GetUserByIdRequest{
		Id: d.Id,
	}

	return &resp
}

func GetUserListRequestToPB(d *GetUserListRequest) *pb.GetUserListRequest {
	if d == nil {
		return nil
	}

	resp := pb.GetUserListRequest{}

	return &resp
}

func PBToGetUserListRequest(d *pb.GetUserListRequest) *GetUserListRequest {
	if d == nil {
		return nil
	}

	resp := GetUserListRequest{}

	return &resp
}

func GetUserListResponseToPB(d *GetUserListResponse) *pb.GetUserListResponse {
	if d == nil {
		return nil
	}

	resp := pb.GetUserListResponse{}

	for _, v := range *d {
		resp.Users = append(resp.Users, UserToPB(&v))
	}

	return &resp
}

func PBToGetUserListResponse(d *pb.GetUserListResponse) *GetUserListResponse {
	if d == nil {
		return nil
	}

	resp := GetUserListResponse{}

	for _, v := range d.Users {
		resp = append(resp, *PBToUser(v))
	}

	return &resp
}

func UpdateUserResponseToPB(d *UpdateUserResponse) *pb.UpdateUserResponse {
	if d == nil {
		return nil
	}

	resp := pb.UpdateUserResponse{}

	return &resp
}

func PBToUpdateUserResponse(d *pb.UpdateUserResponse) *UpdateUserResponse {
	if d == nil {
		return nil
	}

	resp := UpdateUserResponse{}

	return &resp
}

func UserToPB(d *User) *pb.User {
	if d == nil {
		return nil
	}

	resp := pb.User{
		Id:     d.Id,
		Name:   d.Name,
		Active: d.Active,
	}

	return &resp
}

func PBToUser(d *pb.User) *User {
	if d == nil {
		return nil
	}

	resp := User{
		Id:     d.Id,
		Name:   d.Name,
		Active: d.Active,
	}

	return &resp
}
