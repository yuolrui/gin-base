package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuolrui/gin-base/internal/model"
	"github.com/yuolrui/gin-base/internal/service"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var req model.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
