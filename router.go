package main

import (
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/middlewares"
	"time"
)

func RegisterRouter(core *framework.Core) {
	// middleware支持
	core.Use(middlewares.Test3())
	// 静态路由
	core.Get("/", RootControllerHandler)
	core.Get("/foo", FooControllerHandler)
	core.Get("/hello", HelloControllerHandler)
	uGroup := core.Group("/user")
	{
		uGroup.Get("/",UserRootControllerHandler)
		// 动态路由
		uGroup.Get("/list", UserListControllerHandler)
		uGroup.Get("/:id/info",middlewares.TimeoutMiddleware(time.Second),UserInfoPostControllerHandler)
		vGroup := uGroup.Group("/vip")
		{
			vGroup.Get("/version", UserVipVersionControllerHandler)
		}
	}
}
