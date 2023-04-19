package main

import "strings"

func IsRegroup(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 > 5000 || l1 != l2 {
		return false
	}
	for _,v := range s1{
	if strings.Count(s1,string(v))!=strings.Count(s2,string(v)){
		return false
	}
	}
	return true
}