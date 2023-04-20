/*
@author:Deng.l.w
@version:1.20
@date:2023-02-18 10:46
@file:hashMap.go
*/

package main

import "fmt"

// hushMap 建立hasgMap
func hushMap() {
	nameAgeMap := make(map[string]int)
	nameAgeMap["张三"] = 19
	nameAgeMap["李四"] = 23
	nameAgeMap["狂徒"] = 0

	nameAgeMap["狂徒"] = 10

	key1 := "张三"
	value1, ok1 := nameAgeMap[key1] // ok1 为true 时，value1 为对应的值
	if ok1 {
		fmt.Println(key1, value1)
	} else{ 
		fmt.Println(key1, value1)
	}
	for name, age := range nameAgeMap {
		fmt.Println(name, age)
	}
}
