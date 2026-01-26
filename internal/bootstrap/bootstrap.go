package bootstrap

import (
	"log"

	"github.com/yuolrui/gin-base/internal/errno"
	"github.com/yuolrui/gin-base/internal/i18n"
)

func Init() error {
	// 4. 初始化国际化
	i18n.Init()
	if err := i18n.LoadDir("./i18n"); err != nil {
		return err
	}
	if err := i18n.Validate(errno.AllCodes()); err != nil {
		return err
	}
	// 3. 初始化验证翻译器 (重点在此)
	if err := i18n.InitValidator(); err != nil {
		log.Fatalf("初始化校验翻译器失败: %v", err)
	}
	// 后续需要的初始化也可以放这里
	return nil
}
