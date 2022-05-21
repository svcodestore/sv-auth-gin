package auth

import (
	"github.com/svcodestore/sv-auth-gin/service/menu"
	"github.com/svcodestore/sv-auth-gin/service/role"
)

var (
	roleUserService = role.RoleUserService{}
	roleMenuService = role.RoleMenuService{}
	menuService = menu.MenuService{}
)

