package main

import "github.com/afengliz/zett/framework"

func RegisterRouter(core *framework.Core) {
	// 静态路由
	core.Get("/foo", FooControllerHandler)
	core.Get("/hello", HelloControllerHandler)
	core.Post("/hello", HelloPostControllerHandler)
	uGroup := core.Group("/user")
	{
		// 动态路由
		uGroup.Post("/:id",UserInfoPostControllerHandler)
		uGroup.Get("/list", UserListPostControllerHandler)
		vGroup := uGroup.Group("/vip")
		{
			vGroup.Get("/version", UserVipVersionControllerHandler)
		}
	}
}
