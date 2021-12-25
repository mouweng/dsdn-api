package model

import (
	"ginTemplate/config"
	"ginTemplate/model/common"
)

func init() {
	DefaultUserConnection = UserConnection(common.GetDB2DBConnect(config.GetDBConnect(config.DbName)))
}