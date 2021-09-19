package framework

import (
	"net/http"
	"strings"
)

type Core struct{
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {
	// 查询
	getR := make(map[string]ControllerHandler)
	// 新增
	postR := make(map[string]ControllerHandler)
	// 更新
	putR:=make(map[string]ControllerHandler)
	// 删除
	deleteR := make(map[string]ControllerHandler)
	aRouter := make(map[string]map[string]ControllerHandler)
	aRouter[HttPMethodGet] = getR
	aRouter[HttPMethodPost] = postR
	aRouter[HttPMethodPut] = putR
	aRouter[HttPMethodDelete] = deleteR
	return &Core{router: aRouter}
}
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler{
	// 请求方法
	methodName := strings.ToUpper(request.Method)
	// 请求路劲
	rPath := strings.ToUpper(request.URL.Path)
	if mHandlers,ok:=c.router[methodName];ok{
		if handler,ok:= mHandlers[rPath];ok{
			return handler
		}
	}
	return nil
}

func (c *Core) Get(url string,handler ControllerHandler) {
	upUrl := strings.ToUpper(url)
	c.router[HttPMethodGet][upUrl] = handler
}
func (c *Core) Post(url string,handler ControllerHandler){
	upUrl := strings.ToUpper(url)
	c.router[HttPMethodPost][upUrl] = handler
}
func (c *Core) Put(url string,handler ControllerHandler){
	upUrl := strings.ToUpper(url)
	c.router[HttPMethodPut][upUrl] = handler
}
func (c *Core) Delete(url string,handler ControllerHandler){
	upUrl := strings.ToUpper(url)
	c.router[HttPMethodDelete][upUrl] = handler
}

func (c *Core) Group(gName string) IGroup{
	return NewGroup(gName,c)
}




func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request,response)
	router := c.FindRouteByRequest(request)
	if router == nil{
		ctx.Json(http.StatusNotFound,nil)
		return
	}
	ctx.SetHandler(router)
	router(ctx)
}