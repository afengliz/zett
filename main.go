package main

import (
	"context"
	"github.com/afengliz/zett/app/provider/demo"
	"github.com/afengliz/zett/framework/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := gin.New()
	core.Bind(demo.DemoServiceProvider{})
	server := http.Server{Addr: ":8888", Handler: core}
	RegisterRouter(core)
	go func() {
		server.ListenAndServe()
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGQUIT,syscall.SIGTERM)
	<-quit
	if err :=server.Shutdown(context.Background());err != nil{
		log.Fatal("Server Shutdown:",err)
	}
}
