package rpc

import (
	"github.com/svcodestore/sv-auth-gin/rpc/auth"
	"google.golang.org/grpc"
)

func RegisterServer(s *grpc.Server) {
	auth.RegisterAuthRpcServer(s)
}
