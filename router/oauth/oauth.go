package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/api/app"

	"github.com/svcodestore/sv-auth-gin/api"
	"github.com/svcodestore/sv-auth-gin/api/oauth"
)

type OAuthRoutes struct {
}

func (*OAuthRoutes) Init(r *gin.RouterGroup) {
	r.POST("logout", api.Logout)

	oauthG := r.Group("oauth2.0")
	oauthG.POST("/token", oauth.GetAccessToken)

	appG := r.Group("application")
	appG.GET("/current-application", app.GetCurrentApp)
}
