package demo

import (
	"fmt"
	"github.com/afengliz/zett/framework"
)

type DemoServiceProvider struct {

}
var _ framework.ServiceProvider = (*DemoServiceProvider)(nil)

func (d DemoServiceProvider) Name() string {
	return Key
}

func (d DemoServiceProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (d DemoServiceProvider) IsDefer() bool {
	return true
}

func (d DemoServiceProvider) Boot(container framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}

func (d DemoServiceProvider) Register(container framework.Container) framework.NewInstance {
	return NewDemoService
}



