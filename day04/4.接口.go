package main

import "fmt"

//接口是一种类型，抽象的类型
func main() {
	c := cat{}
	c.say()

	var x xiyiji
	h1 := haier{
		name:  "xiaoxuanf",
		price: 100,
	}
	x = h1 //haier实现了洗衣机的全部接口的方法
	x.dry()
	x.wash()

}

type cat struct{}

func (c cat) say() { fmt.Println("miaomiaomiao") }

type sayer interface {
	say()
}

type xiyiji interface {
	wash()
	dry()
}
type haier struct {
	name  string
	price int
}

func (h haier) wash() {
	fmt.Println("xiyifuhaier")
}
func (h haier) dry() {
	fmt.Println("shuaigan")
}
