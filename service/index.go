package service

import (
	"github.com/svcodestore/sv-auth-gin/service/action"
	"github.com/svcodestore/sv-auth-gin/service/app"
	"github.com/svcodestore/sv-auth-gin/service/auth"
	"github.com/svcodestore/sv-auth-gin/service/menu"
	"github.com/svcodestore/sv-auth-gin/service/role"
	"github.com/svcodestore/sv-auth-gin/service/system"
	"github.com/svcodestore/sv-auth-gin/service/user"
)

type Group struct {
	CasbinService         system.CasbinService
	OauthService          system.OauthService
	AppService            app.AppService
	UserService           user.UserService
	MenuService           menu.MenuService
	RoleService           role.RoleService
	RoleUserService       role.RoleUserService
	RoleMenuService       role.RoleMenuService
	RoleUserActionService user.RoleUserActionService
	ActionService         action.ActionService
	ActionMenuService     action.ActionMenuService
	AuthService           auth.AuthService
}

var ServiceGroup = new(Group)
