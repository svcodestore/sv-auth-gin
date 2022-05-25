package core

import (
	"github.com/svcodestore/sv-auth-gin/rpc"
	"log"

	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/initialize"
)

type server interface {
	ListenAndServe() error
}

func commonInit() {
	global.CONFIGURATOR = initialize.InitConfigurator()
	global.LOGGER = initialize.Zap()
	global.DB = initialize.Gorm()
	initialize.DBList()

	initialize.Redis()
}

func RunServer() {
	commonInit()

	db, err := global.DB.DB()
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	routers := initialize.Routers()

	address := global.CONFIG.System.Addr
	s := initServer(address, routers)

	global.LOGGER.Error(s.ListenAndServe().Error())
}

func RunRpcServer() {
	commonInit()

	db, err := global.DB.DB()
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	initialize.InitGrpc(rpc.RegisterServer)
}
