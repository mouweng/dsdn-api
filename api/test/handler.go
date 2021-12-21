package test

import "ginTemplate/api"

func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/test")
		// ping测试链接
		rootRouter.StdGET("ping", Pong)
		// GET Param 解析参数
		rootRouter.StdGET("testParam", DoTestParam)
		// GET Query 解析参数
		rootRouter.StdGET("testQuery", DoTestQuery)
		// 测试Post绑定参数
		rootRouter.StdPOST("testPost", DoTestPost)
	})
}
