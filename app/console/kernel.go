package console

import (
	"github.com/afengliz/zett/app/console/command/demo"
	"github.com/afengliz/zett/framework"
	"github.com/afengliz/zett/framework/cobra"
	"github.com/afengliz/zett/framework/command"
)
// RunCommand 运行root命令
func RunCommand(container framework.Container){
	// 根command
	rootCmd := &cobra.Command{
		// 定义根命令关键字
		Use: "zett",
		// 简短介绍
		Short: "zett 命令",
		// 根命令详细介绍
		Long: "zett 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，" +
			"也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现 cobra 默认的 completion 子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	// 为根 Command 设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架级别的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务级别的命令
	AddAppCommands(rootCmd)
	rootCmd.Execute()
}
// AddAppCommands 绑定业务级别的命令
func AddAppCommands(rootCmd *cobra.Command){
	rootCmd.AddCronCommand("* * * * * *",demo.FooCommand)
}
