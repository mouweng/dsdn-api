package blog

import "ginTemplate/api"

func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/blog")
		// 获取博文
		rootRouter.StdGET("getBlog", DoGetBlog)
		// 添加博文
		rootRouter.StdPOST("addBlog", DoAddBlog)
		// 点赞博文
		rootRouter.StdGET("like", DoLike);
		// 显示是否点赞过
		rootRouter.StdGET("isLike", DoIsLike);
	})
}
