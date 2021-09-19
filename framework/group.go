package framework

import (
	"strings"
)

type IGroup interface {
	Get(url string,handler ControllerHandler)
	Post(url string,handler ControllerHandler)
	Put(url string,handler ControllerHandler)
	Delete(url string,handler ControllerHandler)
	// 支持链式
	Group(gName string) IGroup
}

type group struct {
	segment string
	parent *group
	core *Core
}



func NewGroup(gName string,c *Core) IGroup{
	return &group{segment:strings.ToUpper(gName),core: c}
}

func (c *group) Get(url string,handler ControllerHandler) {
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	c.core.router[HttPMethodGet][upUrl] = handler
}
func (c *group) Post(url string,handler ControllerHandler){
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	c.core.router[HttPMethodPost][upUrl] = handler
}
func (c *group) Put(url string,handler ControllerHandler){
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	c.core.router[HttPMethodPut][upUrl] = handler
}
func (c *group) Delete(url string,handler ControllerHandler){
	upUrl := c.getAbsolutePrefix() + strings.ToUpper(url)
	c.core.router[HttPMethodDelete][upUrl] = handler
}

// 获取前缀路径
func(c *group) getAbsolutePrefix() string{
	if c.parent == nil{
		return c.segment
	}
	return c.parent.getAbsolutePrefix()+c.segment
}

func (c *group) Group(gName string) IGroup {
	return &group{segment: strings.ToUpper(gName),parent: c,core: c.core}
}