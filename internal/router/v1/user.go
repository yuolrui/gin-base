package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yuolrui/gin-base/internal/controller/v1"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/users/:id", v1.GetUser)
	rg.POST("/users", v1.CreateUser)
}
