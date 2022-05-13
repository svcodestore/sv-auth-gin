package service

import (
	"github.com/svcodestore/sv-auth-gin/service/app"
	"github.com/svcodestore/sv-auth-gin/service/menu"
	"github.com/svcodestore/sv-auth-gin/service/system"
	"github.com/svcodestore/sv-auth-gin/service/user"
)

type Group struct {
	CasbinService system.CasbinService
	OauthService  system.OauthService
	AppService    app.AppService
	UserService   user.UserService
	MenuService   menu.MenuService
}

var ServiceGroup = new(Group)
