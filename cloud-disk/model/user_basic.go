package model

import (

	"gorm.io/gorm"
)

// UserBasic is a gorm model
type UserBasic struct {
	gorm.Model
	Name      string 
	Identity  string 
	Password  string 
	Email     string 
	
}

// TableName sets the insert table name for this struct type
func (table UserBasic) TableName() string {
	return "user_basic"
}

