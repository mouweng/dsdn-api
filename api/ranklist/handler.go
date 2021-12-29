package ranklist

import "ginTemplate/api"


func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/ranklist")
		// 获取前n名的点赞数量
		rootRouter.StdGET("likeList", DoLikeList)
	})
}
