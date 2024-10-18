package server

import (
	"context"

	authzextv1 "github.com/grafana/grafana/pkg/services/authz/zanzana/proto/v1"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

var _ authzextv1.AuthzExtentionServiceServer = (*Server)(nil)

func NewAuthz(openfga openfgav1.OpenFGAServiceServer) *Server {
	return &Server{openfga: openfga}
}

type Server struct {
	authzextv1.UnsafeAuthzExtentionServiceServer

	openfga openfgav1.OpenFGAServiceServer
}

// Write implements v1.AuthzExtentionServiceServer.
func (s *Server) Write(context.Context, *authzextv1.WriteRequest) (*authzextv1.WriteResponse, error) {
	panic("unimplemented")
}
