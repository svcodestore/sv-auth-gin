package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
	"strconv"
)

func CreateMenu(c *gin.Context) {
	uid := utils.GetUserID(c)

	pid := c.PostForm("pid")
	applicationId := c.PostForm("applicationId")
	code := c.PostForm("code")
	name := c.PostForm("name")
	sortNo := c.PostForm("sort_no")
	i, e := strconv.Atoi(sortNo)
	var sortNumber uint16
	if e != nil {
		sortNumber = uint16(i)
	}
	path := c.PostForm("path")
	redirect := c.PostForm("redirect")
	component := c.PostForm("component")
	icon := c.PostForm("icon")

	err := menuService.CreateMenu(&model.Menus{
		ID:            "",
		Pid:           pid,
		ApplicationID: applicationId,
		Code:          code,
		Name:          name,
		SortNo:        sortNumber,
		Path:          path,
		Redirect:      redirect,
		Component:     component,
		Icon:          icon,
		CreatedBy:     uid,
		UpdatedBy:     uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}

func DeleteMenuById(c *gin.Context) {
	id := c.Param("id")
	err := menuService.DeleteMenuWithIds(id)
	if err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateMenuById(c *gin.Context) {
	uid := utils.GetUserID(c)

	id := c.PostForm("id")
	pid := c.PostForm("pid")
	applicationId := c.PostForm("applicationId")
	code := c.PostForm("code")
	name := c.PostForm("name")
	sortNo := c.PostForm("sort_no")
	i, e := strconv.Atoi(sortNo)
	var sortNumber uint16
	if e != nil {
		sortNumber = uint16(i)
	}
	path := c.PostForm("path")
	redirect := c.PostForm("redirect")
	component := c.PostForm("component")
	icon := c.PostForm("icon")
	//hide := c.PostForm("hide")
	status := c.PostForm("status")

	updatingMenu := &model.Menus{
		ID:        id,
		UpdatedBy: uid,
	}

	isOnlyUpdateStatus := true

	if pid != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Pid = pid
	}

	if pid != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Pid = pid
	}
	if applicationId != "" {
		isOnlyUpdateStatus = false
		updatingMenu.ApplicationID = applicationId
	}

	if code != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Code = code
	}

	if name != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Name = name
	}

	if sortNo != "" {
		isOnlyUpdateStatus = false
		updatingMenu.SortNo = sortNumber
	}

	if path != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Path = path
	}

	if redirect != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Redirect = redirect
	}

	if component != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Component = component
	}

	if icon != "" {
		isOnlyUpdateStatus = false
		updatingMenu.Icon = icon
	}

	var menu model.Menus
	var err error
	if !isOnlyUpdateStatus {
		err = menuService.UpdateMenuWithId(updatingMenu)
	}
	if err == nil {
		if status == "1" || status == "0" {
			if status == "1" {
				err = menuService.UpdateMenuStatusWithId(true, id, uid)
			} else {
				err = menuService.UpdateMenuStatusWithId(false, id, uid)
			}
		}
		if err == nil {
			response.OkWithData(menu, c)
			return
		}
	}

	response.FailWithMessage(err.Error(), c)
}

func GetAllMenu(c *gin.Context) {
	isAvailable := c.Query("isAvailable")
	appId := c.Query("applicationId")
	if appId != "" {
		menus, err := menuService.MenusWithAppId(appId)
		if err == nil {
			response.OkWithData(menus, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		menus, err := menuService.AllMenu(isAvailable == "1")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(menus, c)
		}
	}
}

func GetMenuById(c *gin.Context) {
	id := c.Param("id")
	menu, err := menuService.MenuWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(menu, c)
	}
}

func BatchCrudMenu(c *gin.Context) {
	id := utils.GetUserID(c)

	var m utils.CrudRequestData
	e := c.ShouldBindJSON(&m)
	if e == nil {
		e = menuService.CrudBatchMenu(id, &m)
		if e == nil {
			response.Ok(c)
			return
		}
	}
	response.FailWithMessage(e.Error(), c)
}
