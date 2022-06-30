package app

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-auth-gin/global"
	"github.com/svcodestore/sv-auth-gin/model/common/response"
	"github.com/svcodestore/sv-auth-gin/service"
	"strconv"
	"strings"
)

var appService = service.ServiceGroup.AppService

func GetCurrentApp(c *gin.Context) {
	id := global.CONFIG.System.Id

	data, err := appService.AppWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		data["clientSecret"] = "***"

		isIntranet := true
		for _, s := range strings.Split(c.Request.Host, ".") {
			_, e := strconv.Atoi(s)
			isIntranet = isIntranet && e == nil
		}

		if data["redirectUris"] != nil {
			redirectUris := strings.Split(data["redirectUris"].(string), "|")
			if len(redirectUris) > 1 {
				if !isIntranet {
					data["redirectUris"] = redirectUris[1]
				}
			}
		}
		if data["loginUris"] != nil {
			loginUris := strings.Split(data["loginUris"].(string), "|")
			if len(loginUris) > 1 {
				if !isIntranet {
					data["loginUris"] = loginUris[1]
				}
			}
		}

		response.OkWithData(data, c)
	}
}

func GetApps(c *gin.Context) {
	apps, err := appService.Apps()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(apps["data"], c)
	}
}
