package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuolrui/gin-base/internal/middleware"
	v1 "github.com/yuolrui/gin-base/internal/router/v1"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(middleware.Lang())
	r.Use(middleware.ErrorHandler())

	// 版本1 API 路由组
	v1Group := r.Group("/api/v1")
	v1.RegisterUserRoutes(v1Group)

	return r
}
