package menu

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type MenuService struct {
}

func (s *MenuService) CreateMenu(m *model.Menus) (err error) {
	m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	err = model.MenusMgr(utils.Gorm()).Create(m).Error

	return
}

func (s *MenuService) DeleteMenuWithIds(ids ...string) (err error) {
	db := global.DB.Where("id IN (?)", ids).Delete(model.Menus{})
	err = db.Error
	return
}

func (s *MenuService) UpdateMenuWithId(m *model.Menus) (err error) {
	id := m.ID
	m.ID = ""
	db := model.MenusMgr(utils.Gorm()).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *MenuService) UpdateMenuStatusWithId(status bool, id, updatedBy string) (err error) {
	err = model.MenusMgr(utils.Gorm()).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *MenuService) AllMenu(isAvailable bool) (menus []*model.Menus, err error) {
	db := utils.Gorm()
	if isAvailable {
		db = db.Where("status = ?", 1)
	}
	menus, err = model.MenusMgr(db).Gets()
	return
}

func (s *MenuService) AllParentMenusWithId(id string) (m []*model.Menus) {
	if menu, e := model.MenusMgr(global.DB).GetFromID(id); e == nil {
		m = append(m, &menu)
		if menu.Pid == "0" {
			return
		}

		m = append(m, s.AllParentMenusWithId(menu.Pid)...)
	}

	return
}

func (s *MenuService) AllSubMenusWithId(id string) (m []*model.Menus) {
	menus := s.SubMenusWithId(id)
	if len(menus) == 0 {
		return
	}
	m = append(m, menus...)

	for _, menu := range menus {
		m = append(m, s.SubMenusWithId(menu.ID)...)
	}

	return
}

func (s *MenuService) ParentMenuWithId(id string) (m model.Menus) {
	if menu, e := model.MenusMgr(global.DB).GetFromID(id); e == nil {
		m, _ = s.MenuWithId(menu.Pid)
	}

	return
}

func (s *MenuService) SubMenusWithId(id string) (m []*model.Menus) {
	m, _ = model.MenusMgr(global.DB).GetFromPid(id)

	return
}

func (s *MenuService) MenuWithId(id string) (menu model.Menus, err error) {
	menu, err = model.MenusMgr(utils.Gorm()).GetFromID(id)
	return
}

func (s *MenuService) MenusWithAppIdAndIds(applicationId string, ids ...string) (menus []*model.Menus, err error) {
	err = model.MenusMgr(global.DB).Where("status = 1 and application_id = ? and id IN (?)", applicationId, ids).Find(&menus).Error
	return
}

func (s *MenuService) AvailableMenus() (menus []*model.Menus, err error) {
	menus, err = s.AllMenu(true)
	return
}

func (s *MenuService) CrudBatchMenu(currentUserId string, data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.Menus
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
		m.CreatedBy = currentUserId
		m.UpdatedBy = currentUserId

		err = s.CreateMenu(&m)

		return
	}, func(b []byte) (err error) {
		var m model.Menus
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}
		m.UpdatedBy = currentUserId

		err = s.UpdateMenuWithId(&m)

		return
	}, func(ids []string) (err error) {
		err = s.DeleteMenuWithIds(ids...)
		return
	})

	return
}

func (s *MenuService) MenusWithAppId(appId string) (menus []*model.Menus, err error) {
	menuMgr := model.MenusMgr(global.DB.Order("pid, sort_no, id"))
	menus, err = menuMgr.GetFromApplicationID(appId)

	return
}
