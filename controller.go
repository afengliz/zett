package main

import "github.com/afengliz/zett/framework"

func FooControllerHandler(ctx *framework.Context) error{
	ctx.GetResponseWriter().Write([]byte("Foo"))
	return nil
}
