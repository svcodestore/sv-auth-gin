package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/api/app"
	"github.com/svcodestore/sv-auth-gin/api/menu"
	"github.com/svcodestore/sv-auth-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(r *gin.RouterGroup) {
	userG := r.Group("user")
	userG.GET("", user.GetUsersByApplicationId)
	userG.GET("current-user", user.GetCurrentUser)

	appsG := r.Group("applications")
	appsG.GET("", app.GetApps)

	menusG := r.Group("menus")
	menusG.GET("", menu.GetAllMenu)
	menuG := r.Group("menu")
	menuG.POST("", menu.CreateMenu)
	menuG.DELETE("/:id", menu.DeleteMenuById)
	menuG.GET("/:id", menu.GetMenuById)
	menuG.PATCH("/:id", menu.UpdateMenuById)
}
