package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
)

func GetAuthMenus(c *gin.Context) {
	applicationId := c.Query("applicationId")
	userId := c.Query("userId")
	if applicationId != "" && userId != "" {
		menus, err := authService.AuthMenusWithApplicationIdAndUserId(applicationId, userId)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(menus, c)
		}
	} else {
		response.Ok(c)
	}
}