package main

import "fmt"

/*
func 函数名（参数）（返回值）{
	函数体
}
*/
func main() {
	sayHi("lc") //函数调用
	res := sum(10, 20, 30, 40)
	fmt.Println(res)
	sum, sub := calc(10, 20)
	fmt.Println(sum, sub)
}

//函数定义,参数
func sayHi(name string) {
	fmt.Println("say hi:", name)
}

//可变参数
func sum(a int, b ...int) int {
	res := a
	for _, v := range b {
		res += v
	}
	return res
}

func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
