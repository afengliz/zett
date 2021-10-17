package command

import (
	"context"
	"github.com/afengliz/zett/framework/cobra"
	"github.com/afengliz/zett/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initAppCommand() *cobra.Command{
	appCommand.AddCommand(appStartCommand)
	return appCommand
}
var appCommand = &cobra.Command{
	Use: "app",
	Short: "业务应用控制命令",
	RunE: func(c *cobra.Command, args []string) error {
		// 打印帮助文档
		c.Help()
		return nil
	},
}
var appStartCommand = &cobra.Command{
	Use: "start",
	Short: "启动一个Web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		kernelService :=container.MustMake(contract.KernelKey).(contract.Kernel)
		engine := kernelService.HttpEngine()
		server := http.Server{Addr: ":8888", Handler: engine}
		go func() {
			server.ListenAndServe()
		}()
		// 当前的goroutine等待信号量
		quit := make(chan os.Signal)
		// 监控信号：SIGINT, SIGTERM, SIGQUIT
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		// 这里会阻塞当前goroutine等待信号
		<-quit
		// 调用Server.Shutdown graceful结束
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		return nil
	},
}