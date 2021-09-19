package main

import "github.com/afengliz/zett/framework"
func RootControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /"))
	return nil
}
func FooControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /foo"))
	return nil
}

func HelloControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /hello"))
	return nil
}



func UserInfoPostControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /user/:id/info"))
	return nil
}

func UserRootControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /user root"))
	return nil
}
func UserListControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /user/list"))
	return nil
}

func UserVipVersionControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("I am /user/vip/version"))
	return nil
}
