package main

import (
	"fmt"
	"strings"
)

func main() {
	f := add()
	res := f(10)
	res = f(10) //20
	fmt.Println(res)

	//生成函数
	f1 := makes(".txt")
	fmt.Println(f1("h1.txt"))
	//调用
	fmt.Println(f1("h1"))

	r1, r2 := cal(100)
	fmt.Println(r1(10), r2(20))
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
		if strings.HasSuffix(s, str) == false {
			return s + str
		}
		return s
	}
}

//后缀生成
func hz(str string) func(string) string { //输入后缀构造新的函数
	return func(s string) string {
		//判断输入是否有后缀
		if strings.HasSuffix(s, str) == false {
			return s + str //没有则加上后缀
		}
		return s //有的话直接返回
	}
}

//函数构造
func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int { //与返回变量一致
		base -= i
		return base
	}
	return add, sub
}
