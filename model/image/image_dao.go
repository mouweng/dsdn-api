package image

import (
	"ginTemplate/config"
	"ginTemplate/model/common"
)

func init() {
	DefaultImageConnection = ImageConnection(common.GetDB2DBConnect(config.GetDBConnect(config.DbName)))
}
