package main

import (
	"fmt"
	"strings"
)

func main() {
	f := add()
	res := f(10)
	res = f(10)
	fmt.Println(res)

	//生成函数
	f1 := makes(".txt")
	fmt.Println(f1("h1.txt"))
	//调用
	fmt.Println(f1("h1"))
}

//闭包返回值,保存x的状态 一起返回组成闭包
func add() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

//闭包返回函数，把s+匿名函数一起 返回
func makes(str string) func(string) string {
	return func(s string) string {
		//如果没有s这个前缀 这返回这个前缀+函数
		if strings.HasPrefix(s, str) == false {
			return s + str
		}
		return s
	}
}
