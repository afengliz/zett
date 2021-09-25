package main

import (
	"context"
	"github.com/afengliz/zett/framework"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()
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
