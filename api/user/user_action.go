package user

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func CreateRoleUserAction(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.RoleUserAction
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.ID = utils.SnowflakeId(utils.RandomInt(1024)).String()
		m.CreatedBy = uid
		m.UpdatedBy = uid
		e = roleUserActionService.CreateRoleUserAction(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func DeleteRoleUserActionById(c *gin.Context) {
	id := c.Param("id")
	err := roleUserActionService.DeleteRoleUserActionWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateRoleUserActionById(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.RoleUserAction
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.UpdatedBy = uid
		e = roleUserActionService.UpdateRoleUserActionWithId(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func GetAllRoleUserAction(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	appId := c.Query("applicationId")
	if appId != "" {
		roleUserActions, err := roleUserActionService.RoleUserActionsWithAppId(appId)
		if err == nil {
			response.OkWithData(roleUserActions, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		roleUserActions, err := roleUserActionService.AllRoleUserAction(isAvailable == "1")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(roleUserActions, c)
		}
	}
}

func GetRoleUserActionById(c *gin.Context) {
	id := c.PostForm("id")
	roleUserAction, err := roleUserActionService.RoleUserActionWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(roleUserAction, c)
	}
}

func BatchCrudRoleUserAction(c *gin.Context) {
	var m utils.CrudRequestData

	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = roleUserActionService.CrudBatchRoleUserAction(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}

	response.FailWithMessage(e.Error(), c)
}
