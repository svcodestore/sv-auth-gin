package app

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/service"
)

var appService = service.ServiceGroup.AppService

func GetCurrentApp(c *gin.Context) {
	id := global.CONFIG.System.Id

	data, err := appService.AppWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		data["clientSecret"] = "***"
		response.OkWithData(data, c)
	}
}

func GetApps(c *gin.Context) {
	apps, err := appService.Apps()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(apps, c)
	}
}
