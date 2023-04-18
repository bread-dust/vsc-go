package validate

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate // validator instance

func init(){
	validate = validator.New() // initialize validator

	validate.SetTagName("v") // set tag name
	
}