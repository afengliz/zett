package main

import "github.com/afengliz/zett/framework"

func RegisterRouter(core *framework.Core){
	core.Get("foo",FooControllerHandler)
}
