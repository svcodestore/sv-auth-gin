package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-auth-gin/middleware"
	"github.com/svcodestore/sv-auth-gin/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	apiGroup := r.Group("api")

	publicApiGroup := apiGroup.Group("")
	router.RouterGroup.OAuth.Init(publicApiGroup)

	privateApiGroup := apiGroup.Group("")
	privateApiGroup.Use(middleware.JWTCheck())
	//privateGroup.Use(middleware.CasbinCheck())
	router.RouterGroup.Api.Init(privateApiGroup)

	return r
}
