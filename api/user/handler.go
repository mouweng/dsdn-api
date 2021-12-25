package user

import "ginTemplate/api"

func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/user")
		// 获取用户
		rootRouter.StdGET("getUser", DoGetUser)
		// 注册用户
		rootRouter.StdPOST("registerUser", DoRegisterUser)

		// 关注用户
		rootRouter.StdGET("followUser", DoFollowUser)

	})
}
