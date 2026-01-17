package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuolrui/gin-base/internal/errno"
	"github.com/yuolrui/gin-base/internal/i18n"
	"github.com/yuolrui/gin-base/internal/middleware"
)

func main() {
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

	r.GET("/user/:id", func(c *gin.Context) {
		c.Error(errno.ErrUserNotExist)
	})

	r.Run(":8080")
}
