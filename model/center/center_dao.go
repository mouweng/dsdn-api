package model

import (
	"ginTemplate/config"
	"ginTemplate/model/common"

)

func init() {
	DefaultCenterConnection = CenterConnection(common.GetDB2DBConnect(config.GetDBConnect(config.DbName)))
}