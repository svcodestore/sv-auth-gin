package initialize

import (
	"gorm.io/gorm"

	"github.com/svcodestore/sv-auth-gin/config"
	"github.com/svcodestore/sv-auth-gin/utils"
)

func Gorm() *gorm.DB {
	return utils.Gorm()
}
func GormMysqlByConfig(m config.DB) *gorm.DB {
	return utils.GormMysqlByConfig(m)
}
