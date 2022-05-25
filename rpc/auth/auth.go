package auth

import (
	"context"
	"github.com/svcodestore/sv-auth-gin/model/rpc/reply"
	"google.golang.org/grpc"

	pb "github.com/svcodestore/sv-auth-gin/proto/auth"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type AuthRpcServer struct {
	pb.UnimplementedAuthServer
}

func RegisterAuthRpcServer(s *grpc.Server) {
	pb.RegisterAuthServer(s, &AuthRpcServer{})
}

func (s *AuthRpcServer) GetAuthMenusWithApplicationIdAndUserId(ctx context.Context, in *pb.GetAuthMenusWithApplicationIdAndUserIdRequest) (*pb.GetAuthMenusWithApplicationIdAndUserIdReply, error) {
	applicationId := in.GetApplicationId()
	userId := in.GetUserId()
	menus, _ := authService.AuthMenusWithApplicationIdAndUserId(applicationId, userId)

	return &pb.GetAuthMenusWithApplicationIdAndUserIdReply{
		AuthMenus: utils.ToRpcStruct(reply.OkWithData(menus)),
	}, nil
}
