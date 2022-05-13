package menu

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type MenuService struct {
}

func (s *MenuService) CreateMenu(m model.Menus) (menu model.Menus, err error) {
	m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	err = model.MenusMgr(utils.Gorm()).Create(m).Error
	if err != nil {
		return
	}
	menu, err = s.MenuWithId(m.ID)
	return
}

func (s *MenuService) DeleteMenuWithId(id string) (isDeleted bool) {
	db := global.DB.Where("id = ?", id).Delete(model.Menus{})
	isDeleted = db.RowsAffected == 1
	return
}

func (s *MenuService) UpdateMenuWithId(m *model.Menus) (menu model.Menus, err error) {
	id := m.ID
	m.ID = ""
	db := model.MenusMgr(utils.Gorm()).Where("id = ?", id).Updates(m)
	if db.RowsAffected == 1 {
		return s.MenuWithId(id)
	}
	err = db.Error

	return
}

func (s *MenuService) UpdateMenuStatusWithId(status bool, id, updatedBy string) (menu model.Menus, err error) {
	err = model.MenusMgr(utils.Gorm()).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error
	if err != nil {
		return
	}
	menu, err = s.MenuWithId(id)
	return
}

func (s *MenuService) AllMenu(isAvailable bool) (menus []*model.Menus, err error) {
	db := utils.Gorm()
	if isAvailable {
		db = db.Where("status = ?", true)
	}
	menus, err = model.MenusMgr(db).Gets()
	return
}

func (s *MenuService) MenuWithId(id string) (menu model.Menus, err error) {
	menu, err = model.MenusMgr(utils.Gorm()).GetFromID(id)
	return
}

func (s *MenuService) AvailableMenus() (menus []*model.Menus, err error) {
	menus, err = s.AllMenu(true)
	return
}
