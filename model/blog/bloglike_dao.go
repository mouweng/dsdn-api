package model

import (
	"ginTemplate/config"
	"ginTemplate/model/common"
)

func init() {
	DefaultBlogLikeConnection = BlogLikeConnection(common.GetDB2DBConnect(config.GetDBConnect(config.DbName)))
}