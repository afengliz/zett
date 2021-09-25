package main

import (
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/middlewares"
	"time"
)

func RegisterRouter(core *framework.Core) {
	// middleware支持
	core.Use(middlewares.Test3(),middlewares.TimeoutMiddleware(time.Second))
	// 静态路由
	core.Get("/", RootControllerHandler)
	core.Get("/foo", FooControllerHandler)
	core.Get("/hello", HelloControllerHandler)
	uGroup := core.Group("/user")
	{
		uGroup.Get("/", UserRootControllerHandler)
		// 动态路由
		uGroup.Get("/list", UserListControllerHandler)
		uGroup.Get("/:id/info",  UserInfoGetControllerHandler)
		vGroup := uGroup.Group("/vip")
		{
			vGroup.Get("/version", UserVipVersionControllerHandler)
		}
		uGroup.Get("/info", GetQueryParamControllerHandler)
		uGroup.Post("/test_form", PostFormParamControllerHandler)
		uGroup.Post("/test_json", TestJsonControllerHandler)
		uGroup.Post("/test_xml", TestXmlControllerHandler)
		uGroup.Post("/test_get_client_address", TestGetClientAddress)
		uGroup.Post("/test_form_file", TestFormFile)
		uGroup.Post("/test_header", TestHeaderControllerHandler)
		uGroup.Post("/test_cookie", TestCookieControllerHandler)
		uGroup.Get("/test_jsonp", TestJsonPControllerHandler)
		uGroup.Post("/test_html", TestHtmlControllerHandler)
		uGroup.Post("/test_text", TestTextControllerHandler)
		uGroup.Get("/test_redirect", TestRedirectControllerHandler)

	}
}
