package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*tree
	middlewares []ControllerHandler
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
func (c *Core) FindRouteByRequest(request *http.Request) *node {
	// 请求方法
	methodName := strings.ToUpper(request.Method)
	// 请求路劲
	rPath := strings.ToUpper(request.URL.Path)
	if mHandlers, ok := c.router[methodName]; ok {
		if mNode := mHandlers.root.matchNode(rPath); mNode != nil {
			return mNode
		}
	}
	return nil
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	upUrl := strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.middlewares...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.router[HttPMethodGet].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}

}
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	upUrl := strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.middlewares...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.router[HttPMethodPost].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	upUrl := strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.middlewares...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.router[HttPMethodPut].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	upUrl := strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.middlewares...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.router[HttPMethodDelete].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = middlewares
}

func (c *Core) Group(gName string) IGroup {
	return NewGroup(gName, c)
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)
	matchNode := c.FindRouteByRequest(request)
	if matchNode == nil {
		ctx.Json(http.StatusNotFound, nil)
		return
	}
	ctx.SetHandler(matchNode.handler...)
	params := matchNode.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "internal error")
		return
	}
}
