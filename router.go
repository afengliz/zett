package main

import "github.com/afengliz/zett/framework"

func RegisterRouter(core *framework.Core) {
	// 静态路由
	core.Get("/", RootControllerHandler)
	core.Get("/foo", FooControllerHandler)
	core.Get("/hello", HelloControllerHandler)
	uGroup := core.Group("/user")
	{
		uGroup.Get("/",UserRootControllerHandler)
		// 动态路由
		uGroup.Get("/list", UserListControllerHandler)
		uGroup.Get("/:id/info",UserInfoPostControllerHandler)
		vGroup := uGroup.Group("/vip")
		{
			vGroup.Get("/version", UserVipVersionControllerHandler)
		}
	}
}
