package user

import "github.com/svcodestore/sv-auth-gin/service"

var (
	userService = service.ServiceGroup.UserService
	roleUserActionService = service.ServiceGroup.RoleUserActionService
)

