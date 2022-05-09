package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(r *gin.RouterGroup) {
	userG := r.Group("user")
	userG.GET("current-user", user.GetCurrentUser)
}
