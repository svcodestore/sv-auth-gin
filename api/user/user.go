package user

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/model/system/request"
	"github.com/svcodestore/sv-auth-gin/service"
)

var userService = service.ServiceGroup.UserService

func GetCurrentUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	id := claims.(*request.CustomClaims).UserId
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
			response.OkWithData(user, c)
		}
	}
	response.Ok(c)
}
