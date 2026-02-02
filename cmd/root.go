package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yuolrui/gin-base/cmd/server" // 替换为你的实际包名
)

var rootCmd = &cobra.Command{
	Use:   "gin-base",
	Short: "A Gin-based clean architecture project",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	// 注册 api server 命令
	rootCmd.AddCommand(server.StartCmd)
}
