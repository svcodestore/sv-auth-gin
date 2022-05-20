package action

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func CreateActionMenu(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.ActionMenu
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.CreatedBy = uid
		m.UpdatedBy = uid
		e = actionMenuService.CreateActionMenu(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func DeleteActionMenuById(c *gin.Context) {
	id := c.Param("id")
	err := actionMenuService.DeleteActionMenuWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateActionMenuById(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.ActionMenu
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.UpdatedBy = uid
		e = actionMenuService.UpdateActionMenuWithId(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func GetAllActionMenu(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	appId := c.Query("applicationId")
	if appId != "" {
		actionMenus, err := actionMenuService.ActionMenusWithAppId(appId)
		if err == nil {
			response.OkWithData(actionMenus, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		actionMenus, err := actionMenuService.AllActionMenu(isAvailable == "1")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(actionMenus, c)
		}
	}
}

func GetActionMenuById(c *gin.Context) {
	id := c.Param("id")
	actionMenus, err := actionMenuService.ActionMenuWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(actionMenus, c)
	}
}

func BatchCrudActionMenu(c *gin.Context) {
	id := utils.GetUserID(c)

	var m utils.CrudRequestData
	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = actionMenuService.CrudBatchActionMenu(id, &m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}
