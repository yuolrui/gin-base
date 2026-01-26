package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuolrui/gin-base/internal/model"
	"github.com/yuolrui/gin-base/internal/response"
	"github.com/yuolrui/gin-base/internal/service"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := service.GetUserByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.Success(user))
}

func CreateUser(c *gin.Context) {
	var req model.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	user, err := service.CreateUser(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.Success(user))
}
