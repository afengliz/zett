package framework

import (
	"log"
	"strings"
)

type IGroup interface {
	Get(url string, handler ControllerHandler)
	Post(url string, handler ControllerHandler)
	Put(url string, handler ControllerHandler)
	Delete(url string, handler ControllerHandler)
	// 支持链式
	Group(gName string) IGroup
}

type group struct {
	segment string
	parent  *group
	core    *Core
}

func NewGroup(gName string, c *Core) IGroup {
	return &group{segment: strings.ToUpper(gName), core: c}
}

func (c *group) Get(url string, handler ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	if err := c.core.router[HttPMethodGet].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *group) Post(url string, handler ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	if err := c.core.router[HttPMethodPost].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *group) Put(url string, handler ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	if err := c.core.router[HttPMethodPut].AddRouter(upUrl, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *group) Delete(url string, handler ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	if err := c.core.router[HttPMethodDelete].AddRouter(upUrl, handler); err != nil {
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

func (c *group) Group(gName string) IGroup {
	return &group{segment: strings.ToUpper(gName), parent: c, core: c.core}
}
