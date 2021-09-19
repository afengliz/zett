package main

import "github.com/afengliz/zett/framework"

func RegisterRouter(core *framework.Core){
	core.Get("/foo",FooControllerHandler)
	core.Get("/hello",HelloControllerHandler)
	core.Post("/hello",HelloPostControllerHandler)
	uGroup := core.Group("/user")
	{
		uGroup.Post("/info",UserInfoPostControllerHandler)
		uGroup.Get("/list",UserListPostControllerHandler)
		vGroup := uGroup.Group("/vip")
		{
			vGroup.Get("/version",UserVipVersionControllerHandler)
		}
	}
}
