package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

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

		// validation错误
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			// 获取对应语言的翻译器
			trans := i18n.GetValidatorTrans(langStr)

			// 翻译错误并取第一个错误（或者拼接所有错误）
			translatedMap := valErrs.Translate(trans)
			var firstMsg string
			for _, v := range translatedMap {
				firstMsg = v
				break // 通常接口返回第一个错误即可
			}

			c.JSON(http.StatusBadRequest, response.Error(errno.ParamErr, firstMsg))
			return
		}

		// 其他自定义错误
		var appErr *errno.AppError
		if errors.As(err, &appErr) {
			msg := i18n.Get(langStr, appErr.Code)
			c.JSON(appErr.HTTPStatus, response.Error(appErr.Code, msg))
			return
		}

		msg := i18n.Get(langStr, errno.InternalErr)
		c.JSON(http.StatusInternalServerError,
			response.Error(errno.InternalErr, msg),
		)
	}
}
