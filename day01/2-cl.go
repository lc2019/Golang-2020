package main

import "fmt"

func main() {
	const PI = 3.14

	//如果批量声明变量某一行没有赋值的时候，默认继承上面一行的值
	const (
		P1 = 100
		P2
		P3
	)

	//iota会在下一个const（被重置为0），没新增1行常量声明const被加1
	const (
		a1 = iota //0
		a2        //1
		a3 = 3    //2
		a4 = iota //3
	)
	fmt.Println(PI)
	fmt.Println(P1, P2, P3)

	fmt.Println(a1, a2, a3, a4)

	const (
		//新增1行iota才加1
		b1, b2 = iota + 1, iota + 2 //1,2
	)

	fmt.Println(b1, b2)
}
