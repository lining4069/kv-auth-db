package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

/*
gin中无法包含所有的验证规则，
对特定的需求需要自定义验证器
*/

// ValidateMobile 验证手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
