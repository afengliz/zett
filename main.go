package main

import (
	"github.com/afengliz/zett/framework"
	"net/http"
)

func main(){
	core := framework.NewCore()
	server := http.Server{Addr: ":8888",Handler: core}
	RegisterRouter(core)
	server.ListenAndServe()
}
