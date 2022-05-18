package role

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func CreateRole(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.Roles
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.CreatedBy = uid
		m.UpdatedBy = uid
		e = roleService.CreateRole(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func DeleteRoleById(c *gin.Context) {
	id := c.Param("id")
	err := roleService.DeleteRoleWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateRoleById(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.Roles
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.UpdatedBy = uid
		e = roleService.UpdateRoleWithId(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func GetAllRole(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	appId := c.Query("applicationId")
	if appId != "" {
		roles, err := roleService.RolesWithAppId(appId)
		if err == nil {
			response.OkWithData(roles, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		roles, err := roleService.AllRole(isAvailable == "1")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(roles, c)
		}
	}
}

func GetRoleById(c *gin.Context) {
	id := c.Param("id")
	role, err := roleService.RoleWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(role, c)
	}
}

func BatchCrudRole(c *gin.Context) {
	id := utils.GetUserID(c)

	var m utils.CrudRequestData
	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = roleService.CrudBatchRole(id, &m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}