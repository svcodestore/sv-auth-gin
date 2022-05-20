package user

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type RoleUserActionService struct {
}

func (s *RoleUserActionService) CreateRoleUserAction(m *model.RoleUserAction) (err error) {
	err = model.RoleUserActionMgr(global.DB).Create(m).Error

	return
}

func (s *RoleUserActionService) DeleteRoleUserActionWithIds(ids ...string) (err error) {
	err = global.DB.Where("id IN (?)", ids).Delete(model.RoleUserAction{}).Error

	return
}

func (s *RoleUserActionService) UpdateRoleUserActionWithId(m *model.RoleUserAction) (err error) {
	id := m.ID
	m.ID = ""
	db := model.RoleUserActionMgr(global.DB).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *RoleUserActionService) UpdateRoleUserActionStatusWithId(status uint8, id, updatedBy string) (err error) {
	err = model.RoleUserActionMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *RoleUserActionService) AllRoleUserAction(isAvailable bool) (userActions []*model.RoleUserAction, err error) {
	db := global.DB
	if isAvailable {
		db = db.Where("status = ?", true)
	}
	userActions, err = model.RoleUserActionMgr(db).Gets()

	return
}

func (s *RoleUserActionService) RoleUserActionWithId(id string) (userAction model.RoleUserAction, err error) {
	userAction, err = model.RoleUserActionMgr(global.DB).GetFromID(id)

	return
}

func (s *RoleUserActionService) RoleUserActionsWithAppId(appId string) (userActions []*model.RoleUserAction, err error) {
	userActions, err = model.RoleUserActionMgr(global.DB).GetFromApplicationID(appId)

	return
}

func (s *RoleUserActionService) AvailableRoleUserAction() (userActions []*model.RoleUserAction, err error) {
	userActions, err = s.AllRoleUserAction(true)

	return
}

func (s *RoleUserActionService) CrudBatchRoleUserAction(currentMenuId string, data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.RoleUserAction
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
		m.CreatedBy = currentMenuId
		m.UpdatedBy = currentMenuId

		err = s.CreateRoleUserAction(&m)

		return
	}, func(b []byte) (err error) {
		var m model.RoleUserAction
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}
		m.UpdatedBy = currentMenuId

		err = s.UpdateRoleUserActionWithId(&m)

		return
	}, func(ids []string) (err error) {
		err = s.DeleteRoleUserActionWithIds(ids...)
		return
	})

	return
}
