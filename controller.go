package main

import (
	"fmt"
	"github.com/afengliz/zett/framework"
	"github.com/spf13/cast"
	"time"
)

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

func UserInfoGetControllerHandler(ctx *framework.Context) error {
	time.Sleep(time.Millisecond * 500)
	paramId, _ := ctx.ParamInt("id", -1)
	ctx.Json(200, fmt.Sprintf("I am /user/:id/info data:%d", paramId))
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

func GetQueryParamControllerHandler(ctx *framework.Context) error {
	userid, _ := ctx.QueryInt("id", -1)
	ctx.GetResponseWriter().Write([]byte(cast.ToString(userid)))
	return nil
}

func PostFormParamControllerHandler(ctx *framework.Context) error {
	userid, _ := ctx.FormInt("id", -1)
	ctx.GetResponseWriter().Write([]byte(cast.ToString(userid)))
	return nil
}