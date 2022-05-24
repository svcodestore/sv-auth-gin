package auth

import (
	"github.com/svcodestore/sv-auth-gin/model"
)

type AuthService struct {
}

func (s *AuthService) AuthMenusWithApplicationAndUserId(applicationId, userId string) (menus []*model.Menus, err error) {
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
	isHasAllMenu := false
	var (
		parentMenus []*model.Menus
		subMenus    []*model.Menus
	)
	for _, menu := range menus {
		if menu.Pid == "0" {
			isHasAllMenu = true
			break
		} else {
			parentMenus = menuService.AllParentMenusWithId(menu.Pid)
			subMenus = menuService.AllSubMenusWithId(menu.ID)
		}
	}
	if isHasAllMenu {
		menus, err = menuService.AllMenu(true)
	}

	menus = append(menus, parentMenus...)
	menus = append(menus, subMenus...)
	var idMaps  = make(map[string]bool)
	var m = make([]*model.Menus, len(idMaps))
	for _, menu := range menus {
		if !idMaps[menu.ID] {
			idMaps[menu.ID] = true
			m = append(m, menu)
		}
	}
	menus = m

	return
}
