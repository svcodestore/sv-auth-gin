package role

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func CreateRoleUser(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.RoleUser
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.CreatedBy = uid
		m.UpdatedBy = uid
		e = roleUserService.CreateRoleUser(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func DeleteRoleUserById(c *gin.Context) {
	id := c.Param("id")
	err := roleUserService.DeleteRoleUserWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateRoleUserById(c *gin.Context) {
	uid := utils.GetUserID(c)

	var m model.RoleUser
	e := c.ShouldBindJSON(&m)
	if e == nil {
		m.UpdatedBy = uid
		e = roleUserService.UpdateRoleUserWithId(&m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}

func GetAllRoleUser(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	roleUsers, err := roleUserService.AllRoleUser(isAvailable == "1")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(roleUsers, c)
	}
}

func GetRoleUserById(c *gin.Context) {
	id := c.Param("id")
	roleUser, err := roleUserService.RoleUserWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(roleUser, c)
	}
}

func BatchCrudRoleUser(c *gin.Context) {
	id := utils.GetUserID(c)

	var m utils.CrudRequestData
	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = roleUserService.CrudBatchRoleUser(id, &m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}