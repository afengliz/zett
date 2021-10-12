package demo

import "github.com/afengliz/zett/framework"

type DemoService struct {
	IDemo
	c framework.Container
}
func (service *DemoService) GetFoo() Foo{
	return Foo{Name: "liyanfeng"}
}
func NewDemoService(params ...interface{}) (interface{},error){
	container := params[0].(framework.Container)
	return &DemoService{c:container},nil
}