package command

import (
	"fmt"
	"github.com/afengliz/zett/framework/cobra"
	"github.com/afengliz/zett/framework/contract"
	"github.com/sevlyar/go-daemon"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//
var cronDaemon bool

func initCronCommand() *cobra.Command {
	cronStartCommand.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "start serve daemon")
	cronCommand.AddCommand(cronStartCommand)
	return cronCommand
}

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取容器
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		pidFolder := appService.RuntimeFolder()
		cronPidFolder := filepath.Join(pidFolder, "cron.pid")
		logFolder := appService.LogFolder()
		cronLogFolder := filepath.Join(logFolder, "cron.log")
		currentFolder := appService.BaseFolder()
		if cronDaemon {
			daemonCtx := daemon.Context{
				LogFileName: cronLogFolder,
				LogFilePerm: 0640,
				PidFileName: cronPidFolder,
				PidFilePerm: 0664,
				WorkDir:     currentFolder,
				Umask:       027,
				// 子进程的参数，按照这个参数设置，子进程的命令为 ./hade cron start --daemon=true
				Args: []string{"", "cron", "start", "--daemon=true"},
			}
			d,err := daemonCtx.Reborn()
			if err != nil{
				return err
			}
			if d != nil{
				// 父进程直接打印启动成功信息，不做任何操作
				fmt.Println("cron serve started, pid:", d.Pid)
				fmt.Println("log file:", cronLogFolder)
				return nil
			}
			// 子进程执行Cron.Run
			defer daemonCtx.Release()
			fmt.Println("daemon started")
			cmd.Root().Cron.Run()
			return nil
		}
		// not deamon mode
		fmt.Println("start cron job")
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := ioutil.WriteFile(cronPidFolder, []byte(content), 0664)
		if err != nil {
			return err
		}
		cmd.Root().Cron.Run()
		return nil
	},
}
