package role

import "github.com/svcodestore/sv-auth-gin/service"

var (
	roleService = service.ServiceGroup.RoleService
	roleUserService = service.ServiceGroup.RoleUserService
	roleMenuService = service.ServiceGroup.RoleMenuService
)
