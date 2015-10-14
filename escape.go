package main

import (
	"strings"
)

var escapeDic = map[string]string{
	",":  "/&#44/",
	"\t": "/&#009/",
	";":  "/&#009/",
	"|":  "/&#124/",
	"\n": "/&#013/",
}

//escape , \n, \t, ;, |
//so that those can be sorted correctly in unix
func Escape(str string) (output string) {
	output = str
	for char, code := range escapeDic {
		output = strings.Replace(output, char, code, -1)
	}
	return output
}

func Unescape(str string) (output string) {
	output = str
	for char, code := range escapeDic {
		output = strings.Replace(output, code, char, -1)
	}
	return output
}
