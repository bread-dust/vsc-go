package main

import (
	"bookstore/pb"
	"context"
	"errors"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaultSize = 10
)
// 使用GORM

func NewDB()(db *gorm.DB,err error){
	dsn := "root:root@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 创建表,自动迁移（把数据表和结构体对应）
	// err = db.AutoMigrate(&Shelf{},&Book{})
	// if err!=nil{
	// 	fmt.Println("sql create error")
	// 	return 
	// }
	return db,err

	// // 创建数据行
	// //u1 := UserInfo{1, "deng", "男", "足球"}
	// //db.Create(&u1)

	// var sf Shelf
	// db.First(&u) //获取第一条记录，主键升序
	// //select * from users order by id limit 1;
	// fmt.Printf("u:%#v\n", u)
	// //更新
	// db.Model(&u).Update("hobby", "篮球")
	// // update user_inofos set hobby='篮球' where id=1

	// //删除
	// db.Delete(&u)
}

// 定义Model
// Shelf 书架
type Shelf struct{
	ID int64 `gorm:"primaryKey"` 
	Theme string
	Size int64
	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct{
	ID int64 `gorm:"primaryKey"`
	Author string
	Title string
	ShelfID int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct{
	db *gorm.DB
}

// CreateShelf 创建书架
func (b *bookstore)CreateShelf(ctx context.Context,data Shelf)(*Shelf,error){
	if len(data.Theme) == 0{
		return nil,errors.New("invalid theme")
	}
	size := data.Size
	if size <=0{
		size = defaultSize
	}
	v := &Shelf{
		Theme:    data.Theme,
		Size:     size,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err:= b.db.WithContext(ctx).Create(&v).Error
	return v,err
}

// GetShelf 获取书架
func (b *bookstore)GetShelf(ctx context.Context,id int64)(*Shelf,error){
	v:=Shelf{}
	err:=b.db.WithContext(ctx).First(&v,id).Error
	return &v,err
}

// ListShelves 书架列表
func (b *bookstore)ListShelves(ctx context.Context)([] *Shelf,error){
	v:=[] *Shelf{}
	err:=b.db.WithContext(ctx).Find(&v).Error
	return v,err
}

// DeleteShelves 删除书架
func (b *bookstore)DeleteShelf(ctx context.Context,id int64)(error){
	err:=b.db.WithContext(ctx).Delete(&Shelf{},id).Error
	return err
}
 

