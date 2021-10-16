package app

import (
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/contract"
)

type ZettAppProvider struct {
	framework.ServiceProvider
	BaseFolder string
}
var _ framework.ServiceProvider = (*ZettAppProvider)(nil)
func (z ZettAppProvider) Name() string {
	return contract.AppKey
}

func (z ZettAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container,z.BaseFolder}
}

func (z ZettAppProvider) IsDefer() bool {
	return false
}

func (z ZettAppProvider) Boot(container framework.Container) error {
	return nil
}

func (z ZettAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewZettApp
}


