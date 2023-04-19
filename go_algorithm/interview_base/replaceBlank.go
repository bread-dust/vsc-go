package main

import (
	"strings"
	"unicode"
)

func ReplaceBlank(s string) string {
	l := len(s)

	if l > 1000 {
		return ""
	}
	for _,v:=range s{
		y := unicode.IsLetter(v)
		if string(v)!=" "&&!y{
			return " "
		}
	}
	return strings.Replace(s," ","%20",-1)
}

