package role

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type RoleUserService struct {
}

func (s *RoleUserService) CreateRoleUser(m *model.RoleUser) (err error) {
	m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	err = model.RoleUserMgr(global.DB).Create(m).Error

	return
}

func (s *RoleUserService) DeleteRoleUserWithIds(ids ...string) (err error) {
	err = global.DB.Where("id IN (?)", ids).Delete(model.RoleUser{}).Error

	return
}

func (s *RoleUserService) UpdateRoleUserWithId(m *model.RoleUser) (err error) {
	id := m.ID
	m.ID = ""
	db := model.RoleUserMgr(global.DB).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *RoleUserService) UpdateRoleUserStatusWithId(status uint8, id, updatedBy string) (err error) {
	err = model.RoleUserMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *RoleUserService) AllRoleUser(isAvailable bool) (roleUsers []*model.RoleUser, err error) {
	db := global.DB
	if isAvailable {
		db = db.Where("status = ?", 1)
	}
	roleUsers, err = model.RoleUserMgr(db).Gets()

	return
}

func (s *RoleUserService) RoleUserWithId(id string) (roleUser model.RoleUser, err error) {
	roleUser, err = model.RoleUserMgr(global.DB).GetFromID(id)

	return
}

func (s *RoleUserService) RoleUsersWithAppId(appId string) (roles []*model.RoleUser, err error) {
	roles, err = model.RoleUserMgr(global.DB).GetFromApplicationID(appId)

	return
}

func (s *RoleUserService) RoleUsersWithAppIdAndUserIds(appId string, userIds ...string) (roles []*model.RoleUser, err error) {
	err = model.RoleUserMgr(global.DB).Where("user_id IN (?) and status = 1 and application_id = ?", userIds, appId).Find(&roles).Error

	return
}

func (s *RoleUserService) RoleUsersWithAppIdAndRoleIds(appId string, roleIds ...string) (roles []*model.RoleUser, err error) {
	roles, err = model.RoleUserMgr(global.DB.Where("role_id IN (?) and status = 1", roleIds)).GetFromApplicationID(appId)

	return
}

func (s *RoleUserService) AvailableRoleUser() (roles []*model.RoleUser, err error) {
	roles, err = s.AllRoleUser(true)

	return
}

func (s *RoleUserService) CrudBatchRoleUser(currentUserId string, data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.RoleUser
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
		m.CreatedBy = currentUserId
		m.UpdatedBy = currentUserId

		err = s.CreateRoleUser(&m)

		return
	}, func(b []byte) (err error) {
		var m model.RoleUser
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}
		m.UpdatedBy = currentUserId

		err = s.UpdateRoleUserWithId(&m)

		return
	}, func(ids []string) (err error) {
		err = s.DeleteRoleUserWithIds(ids...)
		return
	})

	return
}
