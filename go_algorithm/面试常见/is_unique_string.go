package main

import (
	"fmt"
	"strings"
)

// IsUniqueString1 使用strings.Count
func IsUniqueString1(s string)bool{
	if len(s) > 3000{
		return false
	}
	for _,v := range s{
		if v> 127{
			return false 
		}
		v:=string(v)
		if strings.Count(s,v) > 1{
			return false
		}
	}
	return true
}

// IsUniqueString2 使用strings.LastIndex和strings.Index
func IsUniqueString2(s string)bool{
	if len(s) > 3000{
		return false 
	}

	for _,v := range s{
		v := string(v)
		lindex := strings.Index(s,v)
		rindex := strings.LastIndex(s,v)
		return lindex!=rindex
	}
	return true
}

// IsUniqueString3 使用位运算
func IsUniqueString3(s string)bool{
	if len(s) ==0 || len(s) >3000 {
		return false
	}
	mark := make([]uint64,4)
	index:=0
	var offset uint64
	for _,r := range s{
		r := uint64(r)
		fmt.Println(r)
		if r < 64{
			index=0
			offset = 1<<r
		}else if r<128{
			index=1
			offset = 1<<(r-64)
		} else if r<256{
			index=2
			offset = 1<<(r-128)
		}else{
			index=3
			offset = 1<<(r-192)
		}
		fmt.Println(mark[index])
		fmt.Println(offset)
		if mark[index]&offset!=0{
			return false
		}
		mark[index] |= offset
	}
	return true
}