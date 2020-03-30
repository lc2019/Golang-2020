package main

import (
	"fmt"
	"strings"
)

//字符串的操作
func main() {
	path := "/user/mac/go"

	fmt.Println(path)

	//分割
	res := strings.Split(path, "/")
	fmt.Println(res)

	fmt.Println(strings.Join(res, "+"))

}
