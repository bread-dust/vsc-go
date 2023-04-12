package model

import (

	"gorm.io/gorm"
)

// RepositoryPool is a gorm model
type RepositoryPool struct {
	gorm.Model
	Identity string
	Hash     string // 文件hash
	Name     string // 文件名
	Ext      string // 文件后缀
	Size     int64 // 文件大小
	Path     string // 文件路径
	
}

// TableName sets the insert table name for this struct type
func (table RepositoryPool) TableName() string {
	return "repository_pool"
}