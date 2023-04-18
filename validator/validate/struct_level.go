package validate

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type StructLevelUser struct {
	FirstName string
	LastName  string
	Age       uint8  `v:"gte=0,lte=130"`
	Email     string `fld:"e-mail,abc,def" v:"email"`
}

func StructLevel() {
	
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("fld"),",",2)[0]
		if name == "-"{
			return ""
		}
		return name
	})

	validate.RegisterStructValidation(UserStructLevelvValidation,StructLevelUser{})
	user := &StructLevelUser{
		FirstName: "",
		LastName:  "",
		Age:       20,
		Email:     "123",
	}
	err := validate.Struct(user)
	fmt.Println(err)
}

func UserStructLevelvValidation(sl validator.StructLevel){
	user := sl.Current().Interface().(StructLevelUser)
	if len(user.FirstName) == 0 && len(user.LastName) == 0{
		sl.ReportError(user.FirstName,"fname","FirstName","fnameorlname","FirstName or LastName must be set")
		sl.ReportError(user.LastName,"lname","LastName","fnameorlname","FirstName or LastName must be set")}
}
