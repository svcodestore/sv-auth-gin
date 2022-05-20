package action

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type ActionMenuService struct {
}

func (s *ActionMenuService) CreateActionMenu(m *model.ActionMenu) (err error) {
	err = model.ActionMenuMgr(global.DB).Create(m).Error

	return
}

func (s *ActionMenuService) DeleteActionMenuWithIds(ids ...string) (err error) {
	err = global.DB.Where("id IN (?)", ids).Delete(model.ActionMenu{}).Error

	return
}

func (s *ActionMenuService) UpdateActionMenuWithId(m *model.ActionMenu) (err error) {
	id := m.ID
	m.ID = ""
	db := model.ActionMenuMgr(global.DB).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *ActionMenuService) UpdateActionMenuStatusWithId(status uint8, id, updatedBy string) (err error) {
	err = model.ActionMenuMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *ActionMenuService) AllActionMenu(isAvailable bool) (actionMenus []*model.ActionMenu, err error) {
	db := global.DB
	if isAvailable {
		db = db.Where("status = ?", 1)
	}
	actionMenus, err = model.ActionMenuMgr(db).Gets()

	return
}

func (s *ActionMenuService) ActionMenuWithId(id string) (actionMenu model.ActionMenu, err error) {
	actionMenu, err = model.ActionMenuMgr(global.DB).GetFromID(id)

	return
}

func (s *ActionMenuService) ActionMenusWithAppId(appId string) (actionMenus []*model.ActionMenu, err error) {
	actionMenus, err = model.ActionMenuMgr(global.DB.Order("id")).GetFromApplicationID(appId)

	return
}

func (s *ActionMenuService) AvailableActionMenu() (actionMenus []*model.ActionMenu, err error) {
	actionMenus, err = s.AllActionMenu(true)

	return
}

func (s *ActionMenuService) CrudBatchActionMenu(currentMenuId string, data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.ActionMenu
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
		m.CreatedBy = currentMenuId
		m.UpdatedBy = currentMenuId

		err = s.CreateActionMenu(&m)

		return
	}, func(b []byte) (err error) {
		var m model.ActionMenu
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}
		m.UpdatedBy = currentMenuId

		err = s.UpdateActionMenuWithId(&m)

		return
	}, func(ids []string) (err error) {
		err = s.DeleteActionMenuWithIds(ids...)
		return
	})

	return
}
