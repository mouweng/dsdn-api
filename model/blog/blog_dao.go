package model

import (
	"ginTemplate/config"
	"ginTemplate/model/common"
)

func init() {
	DefaultBlogConnection = BlogConnection(common.GetDB2DBConnect(config.GetDBConnect(config.DbName)))
}