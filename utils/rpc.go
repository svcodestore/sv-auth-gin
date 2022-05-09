package utils

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/svcodestore/sv-auth-gin/global"
)

func CallSsoRpcService(cb func(conn *grpc.ClientConn) (reply *structpb.Struct, e error)) (data map[string]interface{}, err error) {
	port := global.CONFIG.System.SsoRpcAddr
	if port == 0 {
		panic("sso server port is 0")
	}
	addr := fmt.Sprintf(":%d", port)

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
