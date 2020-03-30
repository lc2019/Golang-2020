package main

import (
	"fmt"
	"strings"
)

//字符串的操作
func main() {
	path := "/user/mac/golang"

	fmt.Println(path)

	//分割
	res := strings.Split(path, "/")
	fmt.Println(res)

	fmt.Println(strings.Join(res, "+"))

	//字符串拼接
	fmt.Println("hello " + "golang")

	//包含
	fmt.Println(strings.Contains(path, "mac"))

	//前缀后缀
	fmt.Println(strings.HasPrefix(path, "/user"))
	fmt.Println(strings.HasSuffix(path, "/user"))

	//子串的位置
	fmt.Println(strings.Index(path, "u"))
	fmt.Println(strings.LastIndex(path, "g"))

	//join
	s := []string{"hello", "2020"}
	fmt.Println(strings.Join(s, "~"))
}
