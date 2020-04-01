package main

import "fmt"

/*
函数都可以调用
方法只有指定的结构体才能调用
*/
func main() {
	p := &people{"lc", "nan"}
	fmt.Println(p.gender)
	p.dream() //会在里面直接修改为yemen
	fmt.Println(p.gender)

}

type people struct {
	name   string
	gender string
}

//指定接受者的方法，只有people才能调用
func (p *people) dream() {
	p.gender = "yemen"
	fmt.Println("zhengqian ", p.name, p.gender)
}
