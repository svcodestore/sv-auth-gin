package service

import "github.com/svcodestore/sv-auth-gin/service/system"

type Group struct {
	CasbinService system.CasbinService
	OauthService  system.OauthService
}

var ServiceGroup = new(Group)
