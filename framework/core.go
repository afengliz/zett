package framework

import (
	"net/http"
)

type Core struct{
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: make(map[string]ControllerHandler)}
}
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler{
	return c.router["foo"]
}

func (c *Core) Get(pattern string,handler ControllerHandler){
	c.router[pattern] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request,response)
	router := c.FindRouteByRequest(request)
	if router == nil{
		return
	}
	ctx.SetHandler(router)
	router(ctx)
}