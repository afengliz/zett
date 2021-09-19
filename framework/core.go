package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*tree
}

func NewCore() *Core {
	// 查询
	getR := NewTree(HttPMethodGet)
	// 新增
	postR := NewTree(HttPMethodPost)
	// 更新
	putR := NewTree(HttPMethodPut)
	// 删除
	deleteR := NewTree(HttPMethodDelete)
	//
	aRouter := make(map[string]*tree)
	aRouter[HttPMethodGet] = getR
	aRouter[HttPMethodPost] = postR
	aRouter[HttPMethodPut] = putR
	aRouter[HttPMethodDelete] = deleteR
	return &Core{router: aRouter}
}
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// 请求方法
	methodName := strings.ToUpper(request.Method)
	// 请求路劲
	rPath := strings.ToUpper(request.URL.Path)
	if mHandlers, ok := c.router[methodName]; ok {
		if handler := mHandlers.FindHandler(rPath); handler != nil {
			return handler
		}
	}
	return nil
}

func (c *Core) Get(url string, handler ControllerHandler) {
	upUrl := strings.ToUpper(url)
	if err := c.router[HttPMethodGet].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}

}
func (c *Core) Post(url string, handler ControllerHandler) {
	upUrl := strings.ToUpper(url)
	if err := c.router[HttPMethodPost].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Put(url string, handler ControllerHandler) {
	upUrl := strings.ToUpper(url)
	if err := c.router[HttPMethodPut].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Delete(url string, handler ControllerHandler) {
	upUrl := strings.ToUpper(url)
	if err := c.router[HttPMethodDelete].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(gName string) IGroup {
	return NewGroup(gName, c)
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(http.StatusNotFound, nil)
		return
	}
	ctx.SetHandler(router)
	router(ctx)
}
