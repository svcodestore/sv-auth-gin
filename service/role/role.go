package role

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type RoleService struct {
}

func (s *RoleService) CreateRole(m *model.Roles) (err error) {
	err = model.RolesMgr(global.DB).Create(m).Error

	return
}

func (s *RoleService) DeleteRoleWithIds(ids ...string) (err error) {
	err = global.DB.Where("id IN (?)", ids).Delete(model.Roles{}).Error

	return
}

func (s *RoleService) UpdateRoleWithId(m *model.Roles) (err error) {
	id := m.ID
	m.ID = ""
	db := model.RolesMgr(global.DB).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *RoleService) UpdateRoleStatusWithId(status uint8, id, updatedBy string) (err error) {
	err = model.RolesMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *RoleService) AllRole(isAvailable bool) (roles []*model.Roles, err error) {
	db := global.DB
	if isAvailable {
		db = db.Where("status = ?", 1)
	}
	roles, err = model.RolesMgr(db).Gets()

	return
}

func (s *RoleService) RoleWithId(id string) (role model.Roles, err error) {
	role, err = model.RolesMgr(global.DB).GetFromID(id)

	return
}

func (s *RoleService) AvailableRoles() (roles []*model.Roles, err error) {
	roles, err = s.AllRole(true)

	return
}

func (s *RoleService) RolesWithAppId(appId string) (roles []*model.Roles, err error) {
	roles, err = model.RolesMgr(global.DB.Order("pid, id")).GetFromApplicationID(appId)

	return
}

func (s *RoleService) CrudBatchRole(data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.Roles

		err = json.Unmarshal(b, &m)
		if err == nil {
			err = s.CreateRole(&m)
		}

		return
	}, func(b []byte) (err error) {
		var m model.Roles

		err = json.Unmarshal(b, &m)
		if err == nil {
			err = s.UpdateRoleWithId(&m)

		}

		return
	}, func(ids []string) (err error) {
		err = s.DeleteRoleWithIds(ids...)

		return
	})

	return
}
