package main

import "fmt"

/*
函数作用域
*/
var num = 10

func main() {
	test()
	abc := test //函数作为变量传递
	abc()
	fmt.Printf("%p %p", abc, test)
	fmt.Println()

	rs := cals(10, 20, ad) //返回int
	fmt.Println(rs)

	//匿名函数,函数作为返回值
	hi := func() {
		fmt.Println("ni hao yo")
	}
	hi()

	say()
}

func test() {
	//先在内部查找，查找不到在到外部查找
	num := 100
	fmt.Println(num)
}
func ad(x, y int) int {
	return x + y
}

//函数作为类型
func cals(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//匿名函数-匿名函数作为值
func say() {
	func() {
		fmt.Println("ni hao")
	}()
}
