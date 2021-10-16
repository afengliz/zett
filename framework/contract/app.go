package contract

// AppKey 字符串凭证
const AppKey = "hade:app"
// App 定义接口
type App interface {
	// Version 定义当前版本
	Version() string
	// BaseFolder 定义项目基础所在路径
	BaseFolder() string
	// ConfigFolder 定义配置文件所在路径
	ConfigFolder() string
	// LogFolder 定义日志文件所在路径
	LogFolder() string
	// ProviderFolder 定义业务自己的服务提供者所在路径
	ProviderFolder() string
	// MiddlewareFolder 定义业务自己的中间件所在路径
	MiddlewareFolder() string
	// CommandFolder 定义业务自己定义的中间件所在路径
	CommandFolder() string
	// RuntimeFolder 存放运行时的进程 ID 等信息所在路径
	RuntimeFolder() string
	// TestFolder 存放单元测试所在路劲
	TestFolder() string
}
