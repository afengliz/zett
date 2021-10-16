package app

import (
	"errors"
	"flag"
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/contract"
	"github.com/afengliz/zett/framework/utils"
	"path/filepath"
)
var _ contract.App =(*ZettApp)(nil)
type ZettApp struct{
	container framework.Container
	baseFolder string
}

func (z ZettApp) Version() string {
	return "0.0.3"
}

func (z ZettApp) BaseFolder() string {
	if z.baseFolder != "" {
		return z.baseFolder
	}
	var baseFolder string
	flag.StringVar(&baseFolder,"base_folder","","base_folder参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != ""{
		return baseFolder
	}
	return utils.GetExecDirectory()

}

func (z ZettApp) ConfigFolder() string {
	panic("implement me")
}

func (z ZettApp) LogFolder() string {
	return filepath.Join(z.BaseFolder(), "storage","log")
}

func (z ZettApp) ProviderFolder() string {
	return filepath.Join(z.BaseFolder(), "provider")
}

func (z ZettApp) MiddlewareFolder() string {
	return filepath.Join(z.BaseFolder(), "middleware")
}

func (z ZettApp) CommandFolder() string {
	return filepath.Join(z.BaseFolder(), "command")
}

func (z ZettApp) RuntimeFolder() string {
	return filepath.Join(z.BaseFolder(), "storage","runtime")
}

func (z ZettApp) TestFolder() string {
	return filepath.Join(z.BaseFolder(), "test")
}


func NewZettApp(params ...interface{}) (interface{},error){
	if len(params) != 2{
		return nil,errors.New("param error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return ZettApp{container: container,baseFolder: baseFolder},nil
}

