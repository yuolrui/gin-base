package bootstrap

import (
	"fmt"
	"log"

	"github.com/yuolrui/gin-base/internal/config"
	"github.com/yuolrui/gin-base/internal/db" // 假设你创建了 db 工厂包
	"github.com/yuolrui/gin-base/internal/errno"
	"github.com/yuolrui/gin-base/internal/i18n"
	"gorm.io/gorm"
)

// 全局单例，供全项目引用
var (
	DB   *gorm.DB
	Conf *config.Config
)

// Init 应用统一初始化入口
func Init(configPath string) error {
	// 1. 加载配置（由 bootstrap 调度 config 包）
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("load config failed: %w", err)
	}
	Conf = cfg

	// 2. 初始化国际化资源
	if err := initI18n(cfg); err != nil {
		return fmt.Errorf("init i18n failed: %w", err)
	}

	// 3. 初始化 MySQL (GORM)
	if err := initMySQL(&cfg.MySQL); err != nil {
		return fmt.Errorf("init mysql failed: %w", err)
	}

	return nil
}

// 内部拆解：国际化与校验器初始化
func initI18n(cfg *config.Config) error {
	i18n.Init()
	// 路径从配置读取，例如 cfg.App.I18nDir
	if err := i18n.LoadDir("./i18n"); err != nil {
		return err
	}
	if err := i18n.Validate(errno.AllCodes()); err != nil {
		return err
	}
	if err := i18n.InitValidator(); err != nil {
		return err
	}
	return nil
}

// 内部拆解：数据库初始化
func initMySQL(mCfg *config.MySQLConfig) error {
	// 拼接 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		mCfg.User, mCfg.Password, mCfg.Host, mCfg.Port, mCfg.Database, mCfg.Charset)

	// 调用之前建议的 db 包工厂函数
	var err error
	DB, err = db.NewMySQL(dsn, mCfg.MaxIdle, mCfg.MaxOpen)
	if err != nil {
		return err
	}

	log.Println("MySQL initialized successfully")
	return nil
}
