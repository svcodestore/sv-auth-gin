package action

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func CreateAction(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.Actions
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.ID = utils.SnowflakeId(utils.RandomInt(1024)).String()
		m.CreatedBy = uid
		m.UpdatedBy = uid
		e = actionService.CreateAction(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func DeleteActionById(c *gin.Context) {
	id := c.Param("id")
	err := actionService.DeleteActionWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateActionById(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.Actions
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.UpdatedBy = uid
		e = actionService.UpdateActionWithId(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func GetAllAction(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	appId := c.Query("applicationId")
	if appId != "" {
		actions, err := actionService.ActionsWithAppId(appId)
		if err == nil {
			response.OkWithData(actions, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		actions, err := actionService.AllAction(isAvailable == "1")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(actions, c)
		}
	}
}

func GetActionById(c *gin.Context) {
	id := c.Param("id")
	action, err := actionService.ActionWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(action, c)
	}
}

func BatchCrudAction(c *gin.Context) {
	var m utils.CrudRequestData

	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = actionService.CrudBatchAction(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}

	response.FailWithMessage(e.Error(), c)
}
