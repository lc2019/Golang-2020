package main

import "fmt"

func main() {
	defer func() {
		err := recover() //捕获错误后就不在执行
		fmt.Println(err)
	}()
	var a []int
	panic("wo huang le") //pannic后也不会在执行
	a[0] = 10
	fmt.Println("main hanshu")
	fmt.Println(a)
	fmt.Println("main hanshu")
}
