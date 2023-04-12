package model

import (

	"gorm.io/gorm"
)

// UserRepository is a gorm model
type UserRepository struct {
	gorm.Model
	Identity string 
	UserIdentity string
	ParentId     int64  // 父级id
	RepositoryIdentity string 
	Ext      string
	Name     string 

}

// TableName sets the insert table name for this struct type
func (table UserRepository) TableName()string{
	return "user_repository"
}