package main

import "github.com/afengliz/zett/framework"

func FooControllerHandler(ctx *framework.Context) error{
	ctx.GetResponseWriter().Write([]byte("Foo"))
	return nil
}

func HelloControllerHandler(ctx *framework.Context) error{
	ctx.GetResponseWriter().Write([]byte("Hello"))
	return nil
}

func HelloPostControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("Hello Post"))
	return nil
}

func UserInfoPostControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("User Info"))
	return nil
}
func UserListPostControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("User List"))
	return nil
}

func UserVipVersionControllerHandler(ctx *framework.Context) error {
	ctx.GetResponseWriter().Write([]byte("User Vip Version"))
	return nil
}