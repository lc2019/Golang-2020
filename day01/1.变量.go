package main

import "fmt"

func main() {

	//变量的批量声明
	var (
		name = "理想"
		age  = 16
		isOk = true
	)

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(isOk)
	fmt.Printf("name:%T %v\n", isOk, isOk)

	//变量声明的时候赋值
	var s1 string = "雷超"
	fmt.Println(s1)

	//只能声明局部变量
	s2 := "lulu"
	fmt.Println(s2)

	//匿名变量
	a, _ := 10, 20
	fmt.Println(a)

	fmt.Println("人生苦短，let's go")
}
