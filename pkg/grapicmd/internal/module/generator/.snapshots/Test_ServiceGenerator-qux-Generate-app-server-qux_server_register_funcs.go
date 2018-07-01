package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "testapp/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *quxServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterQuxServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *quxServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterQuxServiceHandler(ctx, mux, conn)
}

