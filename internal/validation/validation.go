package validation

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator // 你之前初始化的

func ValidateStruct(c *gin.Context, s interface{}) error {
	langAny, exists := c.Get("lang")
	lang := "zh-CN"
	if exists {
		lang = langAny.(string)
	}

	trans, found := uni.GetTranslator(lang)
	if !found {
		trans, _ = uni.GetTranslator("zh-CN")
	}

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	e := ve[0]

	msg := e.Translate(trans)

	return fmt.Errorf(msg)
}
