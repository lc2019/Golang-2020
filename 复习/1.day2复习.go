package main

import (
	"fmt"
	"strings"
)

func main() {
	//数组的定义
	var a [4]int
	//数组的初始化
	var b = [4]int{1, 2, 3, 4}
	arr := [2]string{"lulu", "leic"}
	//数组的遍历
	for key, value := range arr {
		fmt.Println(key, value)
	}
	fmt.Println(a, b)
	//切片 没有固定长度
	s := []int{1, 2, 3}
	s1 := s
	//切片是引用类型
	s1[0] = 100
	fmt.Println(s, s1)

	//数组是值类型
	c := b
	c[1] = 100
	fmt.Println(b, c)

	//使用copy来实现切片的浅拷贝
	d := make([]int, 3)
	copy(d, s) //copy函数
	d[1] = 200
	fmt.Println(d, s)

	//切片的容量
	fmt.Println(len(s), cap(s))
	fmt.Printf("%p", &s)
	//追加元素
	s = append(s, 1)
	fmt.Println(len(s), cap(s))
	fmt.Printf("%p", &s)

	//切片删除
	d = append(d[:1], d[2:]...) //删除d[1]
	fmt.Println(d)

	//map的声明和初始化
	m1 := make(map[string]int, 5)
	m1["lc"] = 0
	fmt.Println(m1)

	v, ok := m1["lc"] //map中取key的方式
	if ok {
		fmt.Println(" cunz key", v, ok)
	} else {
		fmt.Println("meiy key", v, ok)
	}

	//常见map的用法统计单词次数
	ss := "how do you do are you ok"
	//1.得到分割后的切片
	sl := strings.Split(ss, " ")
	//fmt.Println(sl)
	//2.存储的map
	m2 := make(map[string]int, 0)
	for _, value := range sl {
		//3.判断key是否存在
		v, ok := m2[value]
		if ok {
			//4.存在计算+v 叠加
			m2[value] += v
		} else {
			//5.不存在map的计算值为1
			m2[value] = 1
		}
	}
	fmt.Println(m2)

	//函数参数 返回值
	func(name string) {
		fmt.Println("hello", name)
	}("lc") //匿名函数直接调用

	//得到闭包函数  外部变量+匿名函数
	res := pri("lc")
	//调用闭包
	res("hello")

	//得到闭包函数
	hz := phz(".jpg")
	//调用闭包函数
	str := hz("hello")
	fmt.Println(str)

	//defer

}

//闭包,函数调用外层变量
func pri(name string) func(string) {
	return func(s string) {
		fmt.Println(s, name)
	}
}

//高阶闭包
func phz(name string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, name) == false {
			return s + name
		}
		return s
	}
}
