package validate

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
)

type DbBackedUser struct {
	Name sql.NullString `v:"required"`
	Age sql.NullInt64 `v:"required"`
	Loc *Location `v:"required"`
}

type Location struct{
	X uint
	Y uint 
}	

// CustomField 自定义字段类型验证
func CustomField(){
	// 注册自定义字段类型验证函数，第一个参数是验证函数，第二个参数是需要验证的字段类型
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
	validate.RegisterCustomTypeFunc(ValidLocation,Location{})
	/*	
		if validate.customFuncs == nil {
			validate.customFuncs = make(map[reflect.Type]CustomTypeFunc)
		}
		
		for _,type := range	types{
			validate.customFuncs[reflect.TypeOf(type)] = fn
	*/
	
	// 初始化结构体
	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false},Loc:&Location{X: 1, Y: 1}}

	// 验证结构体
	err := validate.Struct(x)

	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateValuer Name,Age 字段的验证函数 
func ValidateValuer(field reflect.Value) interface{} {
	// 通过反射获取字段的值
	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}

// ValidLocation 验证Location类型
func ValidLocation(field reflect.Value)interface{}{
	if v,ok := field.Interface().(Location);ok{
		if v.X != 0 || v.Y != 0{
			return fmt.Sprintf("%d \n %d",v.X,v.Y)
		}
	}
	return nil
}