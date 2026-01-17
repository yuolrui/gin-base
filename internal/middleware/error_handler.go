package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yuolrui/gin-base/internal/errno"
	"github.com/yuolrui/gin-base/internal/i18n"
	"github.com/yuolrui/gin-base/internal/response"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		lang, _ := c.Get(LangKey)
		langStr, _ := lang.(string)

		var appErr *errno.AppError
		if errors.As(err, &appErr) {
			msg := i18n.Init().Get(langStr, appErr.Code)
			c.JSON(appErr.HTTPStatus, response.Error(appErr.Code, msg))
			return
		}

		msg := i18n.Init().Get(langStr, errno.InternalErr)
		c.JSON(http.StatusInternalServerError,
			response.Error(errno.InternalErr, msg),
		)
	}
}
