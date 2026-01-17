package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuolrui/gin-base/internal/errno"
	"github.com/yuolrui/gin-base/internal/i18n"
	"github.com/yuolrui/gin-base/internal/middleware"
	v1 "github.com/yuolrui/gin-base/internal/router/v1"
)

func InitRouter() *gin.Engine {

	// 1. 加载 i18n
	if err := i18n.LoadDir("./i18n"); err != nil {
		panic(err)
	}

	// 2. 启动校验（示例）
	if err := i18n.Validate([]int{
		errno.OK,
		errno.InvalidParam,
		errno.Unauthorized,
		errno.Forbidden,
		errno.InternalErr,
		errno.UserNotExist,
	}); err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(middleware.Lang())
	r.Use(middleware.ErrorHandler())

	// 版本1 API 路由组
	v1Group := r.Group("/api/v1")
	v1.RegisterUserRoutes(v1Group)

	return r
}
