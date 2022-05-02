package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func Logout(c *gin.Context) {
	t := strings.Split(c.GetHeader("Authorization"), " ")
	if len(t) > 1 {
		token := t[1]
		j := utils.NewJWT()
		claims, _ := j.ParseToken(token)
		affected, err := oauthService.DeleteAccessTokenFromRedis(claims.UserId)
		if err == nil && affected > 0 {
			response.Ok(c)
			return
		}
	}
	response.Fail(c)
}
