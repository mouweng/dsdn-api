package image

import "ginTemplate/api"

func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/image")
		// 获取图片
		rootRouter.StdGET("getImage", DoGetImage)
		// 上传图片
		rootRouter.StdPOST("uploadImage", DoUploadImage)
	})
}
