package app

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/svcodestore/sv-auth-gin/proto/application"
	"github.com/svcodestore/sv-auth-gin/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"time"
)

type AppService struct {
}

func (s AppService) AppWithId(id string) (app gin.H, err error) {
	app, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn) (reply *structpb.Struct, e error) {
		c := pb.NewApplicationClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, e := c.GetApplicationById(ctx, &pb.GetApplicationByIdRequest{
			Id: id,
		})
		if e != nil {
			log.Fatalf("could not get user: %v", err)
		}
		reply = r.GetApplication()
		return
	})
	return
}
