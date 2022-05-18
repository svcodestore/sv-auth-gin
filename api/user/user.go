package user

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func GetCurrentUser(c *gin.Context) {
	id := utils.GetUserID(c)
	user, _ := userService.UserWithId(id)
	response.OkWithData(user, c)
}

func GetUsersByApplicationId(c *gin.Context) {
	id := c.Query("applicationId")
	if id != "" {
		user, err := userService.UsersWithApplicationId(id)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(user["data"], c)
		}
		return
	}
	response.Ok(c)
}
