package oauth

import (
	"encoding/json"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/utils"
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
		clientSecret := global.CONFIG.System.Secret
		code := c.PostForm("code")
		redirectUri := c.PostForm("redirect_uri")
		if code == "" {
			response.UnAuthWithMessage("empty code", c)
			return
		}

		var uri url.URL
		q := uri.Query()

		q.Add("grant_type", grantType)
		q.Add("client_id", clientId)
		q.Add("client_secret", clientSecret)
		q.Add("redirect_uri", redirectUri)
		q.Add("code", code)

		body, err := utils.Post(global.CONFIG.Oauth.TokenUrl+q.Encode(), url.Values{})

		if err != nil {
			response.UnAuthWithMessage("empty code", c)
			return
		}

		var res Result

		json.Unmarshal(body, &res)

		response.OkWithData(res, c)
	}
}
