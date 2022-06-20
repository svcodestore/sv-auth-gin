package role

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func CreateRoleMenu(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.RoleMenu
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.ID = utils.SnowflakeId(utils.RandomInt(1024)).String()
		m.CreatedBy = uid
		m.UpdatedBy = uid
		e = roleMenuService.CreateRoleMenu(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func DeleteRoleMenuById(c *gin.Context) {
	id := c.Param("id")
	err := roleMenuService.DeleteRoleMenuWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateRoleMenuById(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.RoleMenu
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.UpdatedBy = uid
		e = roleMenuService.UpdateRoleMenuWithId(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func GetAllRoleMenu(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	appId := c.Query("applicationId")
	if appId != "" {
		roleMenus, err := roleMenuService.RoleMenusWithAppId(appId)
		if err == nil {
			response.OkWithData(roleMenus, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		roleMenus, err := roleMenuService.AllRoleMenu(isAvailable == "1")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(roleMenus, c)
		}
	}
}

func GetRoleMenuById(c *gin.Context) {
	id := c.PostForm("id")
	roleMenu, err := roleMenuService.RoleMenuWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(roleMenu, c)
	}
}

func BatchCrudRoleMenu(c *gin.Context) {
	var m utils.CrudRequestData

	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = roleMenuService.CrudBatchRoleMenu(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}

	response.FailWithMessage(e.Error(), c)
}
