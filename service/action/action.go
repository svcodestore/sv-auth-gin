package action

import (
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/utils"
)

type ActionService struct {
}

func (s *ActionService) CreateAction(m *model.Actions) (err error) {
	m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
	err = model.ActionsMgr(global.DB).Create(m).Error

	return
}

func (s *ActionService) DeleteActionWithIds(ids ...string) (err error) {
	err = global.DB.Where("id IN (?)", ids).Delete(model.Actions{}).Error

	return
}

func (s *ActionService) UpdateActionWithId(m *model.Actions) (err error) {
	id := m.ID
	m.ID = ""
	db := model.ActionsMgr(global.DB).Where("id = ?", id).Updates(m)
	m.ID = id
	err = db.Error

	return
}

func (s *ActionService) UpdateActionStatusWithId(status uint8, id, updatedBy string) (err error) {
	err = model.ActionsMgr(global.DB).Where("id = ?", id).Select("status").Updates(map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}).Error

	return
}

func (s *ActionService) AllAction(isAvailable bool) (Actions []*model.Actions, err error) {
	db := global.DB
	if isAvailable {
		db = db.Where("status = ?", 1)
	}
	Actions, err = model.ActionsMgr(db).Gets()

	return
}

func (s *ActionService) ActionWithId(id string) (Action model.Actions, err error) {
	Action, err = model.ActionsMgr(global.DB).GetFromID(id)

	return
}

func (s *ActionService) AvailableActions() (Actions []*model.Actions, err error) {
	Actions, err = s.AllAction(true)

	return
}

func (s *ActionService) ActionsWithAppId(appId string) (Actions []*model.Actions, err error) {
	Actions, err = model.ActionsMgr(global.DB.Order("id")).GetFromApplicationID(appId)

	return
}

func (s *ActionService) CrudBatchAction(currentUserId string, data *utils.CrudRequestData) (err error) {
	err = utils.ExecJsonCrudBatch(data, func(b []byte) (err error) {
		var m model.Actions
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		m.ID = utils.SnowflakeId(int64(utils.RandRange(1, 1024))).String()
		m.CreatedBy = currentUserId
		m.UpdatedBy = currentUserId

		err = s.CreateAction(&m)

		return
	}, func(b []byte) (err error) {
		var m model.Actions
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}
		m.UpdatedBy = currentUserId

		err = s.UpdateActionWithId(&m)

		return
	}, func(ids []string) (err error) {
		err = s.DeleteActionWithIds(ids...)
		return
	})

	return
}
