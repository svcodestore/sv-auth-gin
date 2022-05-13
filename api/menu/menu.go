package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"strconv"
)

func CreateMenu(c *gin.Context) {
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

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

	menu, err := menuService.CreateMenu(model.Menus{
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
		response.OkWithData(menu, c)
	}
}

func DeleteMenuById(c *gin.Context) {
	id := c.Param("id")
	isDeleted := menuService.DeleteMenuWithId(id)
	if isDeleted {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateMenuById(c *gin.Context) {
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

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
		ID: id,
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
		menu, err = menuService.UpdateMenuWithId(updatingMenu)
	}
	if err == nil {
		if status == "1" || status == "0" {
			if status == "1" {
				menu, err = menuService.UpdateMenuStatusWithId(true, id, currentUserId)
			} else {
				menu, err = menuService.UpdateMenuStatusWithId(false, id, currentUserId)
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
	menus, err := menuService.AllMenu(isAvailable == "1")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(menus, c)
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