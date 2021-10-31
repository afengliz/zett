package cobra

import (
	"github.com/afengliz/zett/framework"
	"github.com/robfig/cron/v3"
	"log"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}
// CronSpec 保存Cron命令的信息，用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}
func (c *Command) SetParantNull() {
	c.parent = nil
}
// AddCronCommand 是用来创建一个Cron任务的
func (c *Command) AddCronCommand(spec string,cmd *Command){
	root := c.Root()
	if root.Cron == nil{
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}
	// 增加说明信息
	root.CronSpecs = append(root.CronSpecs,CronSpec{
		Type:"normal-cron",
		Cmd: cmd,
		Spec: spec,
	})
	// 拷贝一份 cmd
	var copyCmd Command
	copyCmd = *cmd
	copyCmd.args = []string{}
	copyCmd.SetParantNull()
	copyCmd.SetContainer(root.GetContainer())
	ctx := root.Context()
	// 将要执行的方法添加到root.Cron对象
	root.Cron.AddFunc(spec, func() {
		defer func() {
			if err := recover();err != nil{
				log.Println(err)
			}
		}()
		err := copyCmd.ExecuteContext(ctx)
		if err != nil{
			log.Println(err)
		}
	})
}