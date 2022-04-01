package main

import (
	"github.com/afengliz/zett/app/console"
	"github.com/afengliz/zett/app/http"
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/provider/app"
	"github.com/afengliz/zett/framework/provider/kernel"
)

func main() {
	// 初始化容器
	container := framework.NewZettContainer()
	engine := http.NewHttpEngine()
	container.Bind(&app.ZettAppProvider{})
	container.Bind(&kernel.KernelServiceProvider{
		HttpEngine:engine,
	})
	// 运行root命令
	console.RunCommand(container)
}
