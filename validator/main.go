package main

import (
	"fmt"
	"reflect"
	"validate"
)

func main() {
	validate.StructValidated()
	validate.CustomField()
	validate.StructLevel()
	validate.CustomValidation()

		var a float32 = 1.243
    	reflectType(a)
    	var b int8 = 10
    	reflectType(b)
    
    	var c Cat
    	var d Dog
    	reflectType(c)
    	reflectType(d)


}

func reflectType(x interface{}) {
    	// 我是不知道传入参数的类型
    	obj := reflect.TypeOf(x) // 返回x的动态类型
    	fmt.Println(obj, obj.Name(), obj.Kind())//name 具体信息 ，king 大的种类
    	fmt.Printf("%T", obj)
    }
    
type Cat struct {
    }
    
type Dog struct {
    }
  