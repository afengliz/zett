package main

import (
	"github.com/afengliz/zett/framework/gin"
)

func RegisterRouter(core *gin.Engine) {
	// 静态路由
	core.GET("/", RootControllerHandler)
	core.GET("/foo", FooControllerHandler)
	core.GET("/hello", HelloControllerHandler)
	uGroup := core.Group("/user")
	{
		uGroup.GET("/", UserRootControllerHandler)
		// 动态路由
		uGroup.GET("/list", UserListControllerHandler)
		uGroup.GET("/:id/info",  UserInfoGetControllerHandler)
		vGroup := uGroup.Group("/vip")
		{
			vGroup.GET("/version", UserVipVersionControllerHandler)
		}
		uGroup.GET("/info", GetQueryParamControllerHandler)
		uGroup.POST("/test_form", PostFormParamControllerHandler)
		uGroup.POST("/test_json", TestJsonControllerHandler)
		uGroup.POST("/test_xml", TestXmlControllerHandler)
		uGroup.POST("/test_get_client_address", TestGetClientAddress)
		uGroup.POST("/test_form_file", TestFormFile)
		uGroup.POST("/test_header", TestHeaderControllerHandler)
		uGroup.POST("/test_cookie", TestCookieControllerHandler)
		uGroup.GET("/test_jsonp", TestJsonPControllerHandler)
		uGroup.POST("/test_html", TestHtmlControllerHandler)
		uGroup.POST("/test_text", TestTextControllerHandler)
		uGroup.GET("/test_redirect", TestRedirectControllerHandler)

	}
	dGroup := core.Group("/demo")
	{
		dGroup.POST("/test", PostDemoTestControllerHandler)
	}
}
