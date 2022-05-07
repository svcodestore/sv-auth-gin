package utils

import (
	"fmt"
	"github.com/svcodestore/sv-auth-gin/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func CallSsoRpcService(cb func(conn *grpc.ClientConn)(reply *structpb.Struct, e error)) (data map[string]interface{}, err error) {
	addr := fmt.Sprintf(":%d", global.CONFIG.System.SsoRpcAddr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	reply, e := cb(conn)
	if e != nil {
		return
	}
	data = reply.AsMap()

	return
}