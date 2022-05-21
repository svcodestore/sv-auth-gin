package auth

import (
	"github.com/svcodestore/sv-auth-gin/model"
)

type AuthService struct {
}

func (s *AuthService) AuthMenusWithUserId(applicationId, userId string) (menus []*model.Menus, err error) {
	roleUsers, err := roleUserService.RoleUsersWithAppIdAndUserIds(applicationId, userId)
	if err != nil {
		return
	}
	roleUsersLen := len(roleUsers)
	if roleUsersLen == 0 {
		return
	}
	var roleIds []string
	for i := 0; i < roleUsersLen; i++ {
		roleIds = append(roleIds, roleUsers[i].RoleID)
	}
	roleMenus, err := roleMenuService.RoleMenusWithAppIdAndRoleIds(applicationId, roleIds...)
	roleMenusLen := len(roleMenus)
	if roleMenusLen == 0 {
		return
	}
	var menuIds []string
	for i := 0; i < roleMenusLen; i++ {
		menuIds = append(menuIds, roleMenus[i].MenuID)
	}
	if len(menuIds) == 0 {
		return
	}
	menus, err = menuService.MenusWithAppIdAndIds(applicationId, menuIds...)
	
	return
}
