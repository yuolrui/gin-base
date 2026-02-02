package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	MySQL MySQLConfig `mapstructure:"mysql"`
}

type AppConfig struct {
	Addr string `mapstructure:"addr"`
	Mode string `mapstructure:"mode"`
}
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open"`
}

// LoadConfig 加载并解析配置文件
func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("toml") // 显式指定格式

	// 1. 读取文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file failed: %w", err)
	}

	// 2. 映射到结构体
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %w", err)
	}

	return &cfg, nil
}
