package main

import "fmt"

func main() {
	var x animal
	//值接收
	c := cat1{"buou"}
	x = c
	x.speak()
	x.move()
	fmt.Println(x) //{buou}

	//指针接收
	tom := &cat1{"tom"}
	x = tom
	fmt.Println(x, tom) //tom
}

type animal interface {
	speak()
	move()
}
type cat1 struct {
	name string
}

func (c cat1) speak() {
	fmt.Println("miaomiao")
}
func (c cat1) move() {
	fmt.Println("zoumaobu")
}
