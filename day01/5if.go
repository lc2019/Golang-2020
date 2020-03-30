package main

import "fmt"

func main() {
	//变量的作用域在{}里面
	age := 31
	if age > 18 { //如果条件成立执行
		fmt.Println("gogogo")
	} else { //否则执行
		fmt.Println("goon 123")
	}

	//多个条件判断
	if age > 35 {
		fmt.Println("jiayou")
	} else if age > 28 {
		fmt.Println("fendou")
	} else {
		fmt.Println("go to")
	}
}
