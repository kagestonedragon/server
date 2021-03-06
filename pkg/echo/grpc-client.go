// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package echo

import (
	"context"
	"errors"

	pb "github.com/kagestonedragon/server/internal/echopb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// NewGRPCClient returns an Service backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func NewGRPCClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) Service {
	// global client middlewares
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	}

	return endpoints{
		// Each individual endpoint is an grpc/transport.Client (which implements
		// endpoint.Endpoint) that gets wrapped with various middlewares. If you
		// made your own client library, you'd do this work there, so your server
		// could rely on a consistent set of client behavior.
		GetEchoEndpoint: grpctransport.NewClient(
			conn,
			"echopb.EchoService",
			"GetEcho",
			encodeGRPCGetEchoListRequest,
			decodeGRPCGetEchoListResponse,
			pb.GetEchoListResponse{},
			options...,
		).Endpoint(),
	}
}

func encodeGRPCGetEchoListRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*GetEchoListRequest)
	if !ok {
		return nil, errors.New("encodeGRPCGetEchoListRequest wrong request")
	}

	return GetEchoListRequestToPB(inReq), nil
}

func decodeGRPCGetEchoListResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*pb.GetEchoListResponse)
	if !ok {
		return nil, errors.New("decodeGRPCGetEchoListResponse wrong response")
	}

	resp := PBToGetEchoListResponse(inResp)

	return *resp, nil
}
