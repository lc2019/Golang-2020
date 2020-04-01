package main

import (
	"fmt"
)

type stu2 struct {
	name string
	age  int
}

//构造函数
func newStu2(name string, age int) *stu2 {
	return &stu2{
		name: name,
		age:  age,
	}
}

func main() {
	//实例化
	s2 := stu2{"lc", 18}
	fmt.Println(s2.name, s2.age)

	//new产生的是指针
	s3 := new(stu2)
	//(*stu2).name
	s3.age = 18
	s3.name = "lc"
	fmt.Println(*s3)

	//实例化
	s4 := &stu2{"lc", 18}
	fmt.Println(s4.name, s4.age)

	//结构体内存布局
	type test struct {
		a int8
		b int8
		c int8
	}
	te := test{1, 2, 3}
	fmt.Println(&(te.a), &(te.b), &(te.c))

	//结构体 值赋值
	test2 := te
	test2.a = 10
	fmt.Println(te, test2)

	//结构体指针赋值
	test3 := &te
	test3.a = 100
	fmt.Println(te, test3)

	ss := newStu2("lulu", 18)
	fmt.Println(ss.age, ss.name)
}
