package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomValidationUser struct {
	Name   string `v:"required,is-deng"`
	Gender int    `v:"gender"`
}

// CustomValidation 自定义标签验证
func CustomValidation() {
	// 注册自定义标签，第一个参数是标签名，第二个参数是验证函数
	validate.RegisterValidation("is-deng", ValidateName)
	validate.RegisterValidation("gender", ValidateGender)
	user := &CustomValidationUser{
		Name:   "deng",
		Gender: 4,
}
	err:= validate.Struct(user)
	fmt.Println(err)
}

// ValidateName 自定义"is-deng"标签验证函数
func ValidateName(fl validator.FieldLevel) bool {
		return fl.Field().String() == "deng"
}

// ValidateGender 自定义"gender"标签验证函数
func ValidateGender(fl validator.FieldLevel) bool {
	i := fl.Field().Int()
	if i!=0 && i !=1 && i!=2{
		return false
	}
	return true
}