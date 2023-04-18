package validate

import (
	"fmt"
)

type User struct {
	Name string `v:"required,alpha"` // required,alpha
	Age uint8 `v:"number,min=18,max=100"`
	Phone string `v:"required_without=Email,omitempty"` // required_without:Email
	Email string `v:"required_without=Phone,omitempty"` // required_without:Phone
	Hobby []string  `v:"required,min=2,dive,required,alphaunicode,min=3"`
}



func StructValidated(){
	user:= &User{
		Name:  "deng",
		Age:   20,
		Phone: "+861234567890",
		Email: "",
		Hobby: []string{"邓丽伟","eff"},
	}
	err := validate.Struct(user)
	fmt.Println(err)
}

