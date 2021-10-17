package kernel

import (
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/contract"
	"github.com/afengliz/zett/framework/gin"
)

type KernelServiceProvider struct {
	HttpEngine *gin.Engine
}
var _ framework.ServiceProvider = (*KernelServiceProvider)(nil)

func (k KernelServiceProvider) Name() string {
	return contract.KernelKey
}

func (k KernelServiceProvider) Params(container framework.Container) []interface{} {
	return []interface{}{k.HttpEngine}
}

func (k KernelServiceProvider) IsDefer() bool {
	return false
}

func (k KernelServiceProvider) Boot(container framework.Container) error {
	if k.HttpEngine == nil{
		k.HttpEngine = gin.Default()
	}
	k.HttpEngine.SetContainer(container)
	return nil
}

func (k KernelServiceProvider) Register(container framework.Container) framework.NewInstance {
	return NewKernelService
}

