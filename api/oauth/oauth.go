package oauth

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/service"
)

var (
	oauthService = service.ServiceGroup.OauthService
	appService   = service.ServiceGroup.AppService
)

type Token struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         interface{} `json:"user"`
}

type Result struct {
	Code    int    `json:"code"`
	Data    Token  `json:"data"`
	Message string `json:"message"`
}

func GetAccessToken(c *gin.Context) {
	grantType := c.PostForm("grant_type")

	if grantType == "authorization_code" {
		clientId := c.PostForm("client_id")
		code := c.PostForm("code")
		redirectUri := c.PostForm("redirect_uri")
		if code == "" {
			response.UnAuthWithMessage("empty code", c)
			return
		}

		clientSecret, err := appService.AppClientSecretWithClientId(clientId)
		if err != nil {
			response.UnAuthWithMessage(err.Error(), c)
			return
		}

		resp, err := oauthService.GetAccessToken(grantType, clientId, clientSecret, code, redirectUri)
		if err == nil {
			response.OkWithData(resp, c)
		}
	}
}
