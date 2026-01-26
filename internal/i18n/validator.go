package i18n

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni   *ut.UniversalTranslator
	trans = make(map[string]ut.Translator)
)

// InitValidator 初始化校验翻译器
func InitValidator() error {
	zhT := zh.New()
	enT := en.New()
	uni = ut.New(enT, zhT, enT)

	// 注册支持的语言
	locales := []string{"zh", "en"}
	for _, locale := range locales {
		t, ok := uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 获取验证引擎
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if !ok {
			continue
		}

		// 注册对应语言的默认翻译逻辑
		switch locale {
		case "zh":
			_ = zhtranslations.RegisterDefaultTranslations(v, t)
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(v, t)
		}

		// 重点：注册自定义标签获取逻辑，优先使用 json 标签作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		trans[locale] = t
	}
	return nil
}

// GetValidatorTrans 根据语言代码获取翻译器
func GetValidatorTrans(lang string) ut.Translator {
	// 简单清洗 lang 字符串，处理 zh-CN 或 en-US 这种格式
	lang = strings.ToLower(lang)
	if strings.Contains(lang, "zh") {
		lang = "zh"
	} else if strings.Contains(lang, "en") {
		lang = "en"
	} else {
		lang = "zh" // 默认语言
	}

	if t, ok := trans[lang]; ok {
		return t
	}
	return trans["zh"]
}
