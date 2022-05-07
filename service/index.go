package service

import (
	"github.com/svcodestore/sv-auth-gin/service/app"
	"github.com/svcodestore/sv-auth-gin/service/system"
)

type Group struct {
	CasbinService system.CasbinService
	OauthService  system.OauthService
	AppService    app.AppService
}

var ServiceGroup = new(Group)
