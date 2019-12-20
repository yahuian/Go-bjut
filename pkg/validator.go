package pkg

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var trans ut.Translator

// 初始化翻译器
func InitTranslator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("binding.Validator.Engine() is not the type *validator.Validate")
	}
	chinese := zh.New()
	uni := ut.New(chinese, chinese)
	trans, _ = uni.GetTranslator("zh")
	if err := zhTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return err
	}
	return nil
}

// 用于将validator.v9返回的错误信息，翻译为一种用户友好的提示
func Warn(err error) []string {
	errorsMap, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}
	var result []string
	for _, v := range errorsMap {
		result = append(result, v.Translate(trans))
	}
	return result
}
