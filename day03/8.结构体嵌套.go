package main

import "fmt"

/*
匿名字段只有1个
*/
func main() {
	t := test{name: "lc"}
	fmt.Println(t.name, t.string)

	s1 := stu{"lc", address{"hu", "tm"}, email{"lcc@"}}
	//匿名结构体字段冲突的时候不支持简写
	fmt.Println(s1.name, s1.address.df, s1.email.df)
	//匿名结构体结构体字段唯一的时候，直接递归查找到address
}

type test struct {
	name   string
	string //匿名结构体变量 只能有一个
}
type stu struct {
	name    string
	address //结构体嵌套 ，匿名结构体嵌套，如果有2个匿名结构体必须指定指端
	email
}

type address struct {
	cs string
	df string
}

type email struct {
	df string
}
