package role

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type RoleMenuService struct {
}

func (s *RoleMenuService) CreateRoleMenu(m *model.RoleMenu) (err error) {
	err = model.RoleMenuMgr(global.DB).Create(m).Error

	return
}

func (s *RoleMenuService) DeleteRoleMenuWithIds(ids ...string) (err error) {
	err = global.DB.Where("id IN (?)", ids).Delete(model.RoleMenu{}).Error

	return
}

func (s *RoleMenuService) UpdateRoleMenuWithId(m *model.RoleMenu) (err error) {
	id := m.ID
	m.ID = ""
	db := model.RoleMenuMgr(global.DB).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *RoleMenuService) UpdateRoleMenuStatusWithId(status uint8, id, updatedBy string) (err error) {
	err = model.RoleMenuMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *RoleMenuService) AllRoleMenu(isAvailable bool) (roleMenus []*model.RoleMenu, err error) {
	db := global.DB
	if isAvailable {
		db = db.Where("status = ?", 1)
	}
	roleMenus, err = model.RoleMenuMgr(db).Gets()

	return
}

func (s *RoleMenuService) RoleMenuWithId(id string) (roleMenu model.RoleMenu, err error) {
	roleMenu, err = model.RoleMenuMgr(global.DB).GetFromID(id)

	return
}

func (s *RoleMenuService) RoleMenusWithAppId(appId string) (roleMenus []*model.RoleMenu, err error) {
	roleMenus, err = model.RoleMenuMgr(global.DB).GetFromApplicationID(appId)

	return
}

func (s *RoleMenuService) AvailableRoleMenu() (roleMenus []*model.RoleMenu, err error) {
	roleMenus, err = s.AllRoleMenu(true)

	return
}

func (s *RoleMenuService) CrudBatchRoleMenu(currentMenuId string, data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.RoleMenu
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
		m.CreatedBy = currentMenuId
		m.UpdatedBy = currentMenuId

		err = s.CreateRoleMenu(&m)

		return
	}, func(b []byte) (err error) {
		var m model.RoleMenu
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}
		m.UpdatedBy = currentMenuId

		err = s.UpdateRoleMenuWithId(&m)

		return
	}, func(ids []string) (err error) {
		err = s.DeleteRoleMenuWithIds(ids...)
		return
	})

	return
}
