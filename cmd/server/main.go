package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/yuolrui/gin-base/internal/bootstrap"
	"github.com/yuolrui/gin-base/internal/router"
)

var (
	configPath string
	StartCmd   = &cobra.Command{
		Use:          "server",
		Short:        "Start the API server",
		Example:      "gin-base server -c configs/config.toml",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// 1. 在运行前执行初始化
			return bootstrap.Init(configPath)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// 2. 启动服务
			return run()
		},
	}
)

func init() {
	// 定义命令行参数：-c 或 --config
	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "configs/config.toml", "config file path")
}

func run() error {
	// 设置 Gin 模式
	if bootstrap.Conf.App.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化路由
	r := router.InitRouter()

	fmt.Printf("\n  Server running at:\n")
	fmt.Printf("  - Local:   http://localhost%s\n", bootstrap.Conf.App.Addr)

	// 启动服务
	return r.Run(bootstrap.Conf.App.Addr)
}
