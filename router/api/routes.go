package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/api/action"
	"github.com/svcodestore/sv-auth-gin/api/app"
	"github.com/svcodestore/sv-auth-gin/api/auth"
	"github.com/svcodestore/sv-auth-gin/api/menu"
	"github.com/svcodestore/sv-auth-gin/api/role"
	"github.com/svcodestore/sv-auth-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(r *gin.RouterGroup) {
	usersG := r.Group("users")
	usersG.GET("", user.GetUsersByApplicationId)
	userG := r.Group("user")
	userG.GET("current-user", user.GetCurrentUser)

	userActionsG := r.Group("user-actions")
	userActionsG.GET("", user.GetAllRoleUserAction)
	userActionsG.POST("/batch", user.BatchCrudRoleUserAction)
	userActionG := r.Group("user-action")
	userActionG.POST("", user.CreateRoleUserAction)
	userActionG.DELETE("/:id", user.DeleteRoleUserActionById)
	userActionG.GET("/:id", user.GetRoleUserActionById)
	userActionG.PATCH("/:id", user.UpdateRoleUserActionById)

	appsG := r.Group("applications")
	appsG.GET("", app.GetApps)

	menusG := r.Group("menus")
	menusG.GET("", menu.GetAllMenu)
	menusG.POST("/batch", menu.BatchCrudMenu)
	menuG := r.Group("menu")
	menuG.POST("", menu.CreateMenu)
	menuG.DELETE("/:id", menu.DeleteMenuById)
	menuG.GET("/:id", menu.GetMenuById)
	menuG.PATCH("/:id", menu.UpdateMenuById)

	actionsG := r.Group("actions")
	actionsG.GET("", action.GetAllAction)
	actionsG.POST("/batch", action.BatchCrudAction)
	actionG := r.Group("action")
	actionG.POST("", action.CreateAction)
	actionG.DELETE("/:id", action.DeleteActionById)
	actionG.GET("/:id", action.GetActionById)
	actionG.PATCH("/:id", action.UpdateActionById)

	actionMenusG := r.Group("action-menus")
	actionMenusG.GET("", action.GetAllActionMenu)
	actionMenusG.POST("/batch", action.BatchCrudActionMenu)
	actionMenuG := r.Group("action-menu")
	actionMenuG.POST("", action.CreateActionMenu)
	actionMenuG.DELETE("/:id", action.DeleteActionMenuById)
	actionMenuG.GET("/:id", action.GetActionMenuById)
	actionMenuG.PATCH("/:id", action.UpdateActionMenuById)

	rolesG := r.Group("roles")
	rolesG.GET("", role.GetAllRole)
	rolesG.POST("/batch", role.BatchCrudRole)
	roleG := r.Group("role")
	roleG.POST("", role.CreateRole)
	roleG.DELETE("/:id", role.DeleteRoleById)
	roleG.GET("/:id", role.GetRoleById)
	roleG.PATCH("/:id", role.UpdateRoleById)

	roleMenusG := r.Group("role-menus")
	roleMenusG.GET("", role.GetAllRoleMenu)
	roleMenusG.POST("/batch", role.BatchCrudRoleMenu)
	roleMenuG := r.Group("role-menu")
	roleMenuG.POST("", role.CreateRoleMenu)
	roleMenuG.DELETE("/:id", role.DeleteRoleMenuById)
	roleMenuG.GET("/:id", role.GetRoleMenuById)
	roleMenuG.PATCH("/:id", role.UpdateRoleMenuById)

	roleUsersG := r.Group("role-users")
	roleUsersG.GET("", role.GetAllRoleUser)
	roleUsersG.POST("/batch", role.BatchCrudRoleUser)
	roleUserG := r.Group("role-user")
	roleUserG.POST("", role.CreateRoleUser)
	roleUserG.DELETE("/:id", role.DeleteRoleUserById)
	roleUserG.GET("/:id", role.GetRoleUserById)
	roleUserG.PATCH("/:id", role.UpdateRoleUserById)

	authG := r.Group("authorization")
	authG.GET("/user-menus", auth.GetAuthMenus)
}
