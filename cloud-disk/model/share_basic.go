package model

import "gorm.io/gorm"

// ShareBasic is a gorm model
type ShareBasic struct {

	gorm.Model
	Identity               string
	UserIdentity           string
	RepositoryIdentity     string
	UserRepositoryIdentity string
	ExpiredTime            int	// 过期时间
	ClickNum               int // 点击次数
}

// TableName sets the insert table name for this struct type
func (table ShareBasic) TableName() string {
	return "share_basic"
}