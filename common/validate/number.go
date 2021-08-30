package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
)

// 验证是否是正整数
func ValidateUintFormat(fl validator.FieldLevel) bool {
	// todo 目前只支持 字段为 string 的配置 以后再解决
	val := fl.Field().String()
	matched, err := regexp.MatchString(`^\d+$`, val)
	if err != nil {
		return false
	}
	if !matched {
		return false
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return i > 0
}
