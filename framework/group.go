package framework

import (
	"log"
	"strings"
)

type IGroup interface {
	Get(url string, handlers ...ControllerHandler)
	Post(url string, handlers ...ControllerHandler)
	Put(url string, handlers ...ControllerHandler)
	Delete(url string, handlers ...ControllerHandler)
	// 支持链式
	Group(gName string) IGroup
	Use(middlewares ...ControllerHandler)
}

type group struct {
	segment     string
	parent      *group
	core        *Core
	middlewares []ControllerHandler
}

func NewGroup(gName string, c *Core) IGroup {
	return &group{
		segment: strings.ToUpper(gName),
		core:    c,
	}
}
func (c *group) Use(middlewares ...ControllerHandler) {
	c.middlewares = middlewares
}
func (c *group) Get(url string, handlers ...ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.getAllMiddlewares()...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.core.router[HttPMethodGet].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *group) Post(url string, handlers ...ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.getAllMiddlewares()...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.core.router[HttPMethodPost].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *group) Put(url string, handlers ...ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.getAllMiddlewares()...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.core.router[HttPMethodPut].AddRouter(upUrl, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *group) Delete(url string, handlers ...ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	var allHandlers []ControllerHandler
	allHandlers = append(allHandlers, c.getAllMiddlewares()...)
	allHandlers = append(allHandlers, handlers...)
	if err := c.core.router[HttPMethodDelete].AddRouter(upUrl, handlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 获取前缀路径
func (c *group) getAbsolutePrefix() string {
	if c.parent == nil {
		return c.segment
	}
	return c.parent.getAbsolutePrefix() + c.segment
}
func (c *group) getAllMiddlewares() []ControllerHandler {
	var tmp []ControllerHandler
	tmp = append(tmp, c.core.middlewares...)
	tmp = append(tmp, c.getGroupMiddlewares()...)
	return tmp
}

func (c *group) getGroupMiddlewares() []ControllerHandler {
	if c.parent == nil {
		return c.middlewares
	}
	var tmp []ControllerHandler
	tmp = append(tmp, c.parent.getGroupMiddlewares()...)
	tmp = append(tmp, c.middlewares...)
	return tmp
}

func (c *group) Group(gName string) IGroup {
	return &group{
		segment: strings.ToUpper(gName),
		parent:  c,
		core:    c.core,
	}
}
