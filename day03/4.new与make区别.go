package main

import "fmt"

func main() {
	s1 := new([]int)
	fmt.Println(s1) //返回的是一个指针,没有初始化不能存储

	s2 := make([]int, 3)
	fmt.Println(s2) //返回的直接是切片，已经初始化

	*s1 = make([]int, 5)
	(*s1)[0] = 10
	fmt.Println(s1)

}
