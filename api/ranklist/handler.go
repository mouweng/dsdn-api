package ranklist

import "ginTemplate/api"

func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/ranklist")
		// 获取用户
		rootRouter.StdGET("test", DoTest)
	})
}
