package message

import "ginTemplate/api"


func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/message")
		//查询
		rootRouter.StdGET("getCenter", DoGetCenter)
		// 添加
		rootRouter.StdPOST("addCenter", DoAddCenter)
		// 删除（事务的运用）
		rootRouter.StdGET("delCenter", DoDelCenter)
	})
}
