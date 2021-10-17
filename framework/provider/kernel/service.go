package kernel

import (
	"github.com/afengliz/zett/framework/contract"
	"github.com/afengliz/zett/framework/gin"
	"net/http"
)

func NewKernelService(params ...interface{}) (interface{},error){
	engine := params[0].(*gin.Engine)
	return &ZettKernelService{engine:engine},nil
}

type ZettKernelService struct {
	engine *gin.Engine
}


var _ contract.Kernel = (*ZettKernelService)(nil)

func (z ZettKernelService) HttpEngine() http.Handler {
	return z.engine
}