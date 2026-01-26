package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	LangKey = "lang"
)

func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("X-Lang")
		if lang == "" {
			lang = c.GetHeader("Accept-Language")
		}
		if lang == "" {
			lang = "zh-CN"
		}
		c.Set(LangKey, lang)
		c.Next()
	}
}
