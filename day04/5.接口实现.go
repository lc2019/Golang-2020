package main

import "fmt"

func main() {
	var x annimal
	c := Cat{"buou"}
	x = c
	x.move()
	x.speak()
}

type speaker interface {
	speak()
}

type mover interface {
	move()
}
type annimal interface {
	speaker
	mover //接口嵌套
}
type Cat struct {
	name string
}

//接口实现
func (c Cat) speak() {
	fmt.Println("miaomiaomiao")
}

//接口实现
func (c Cat) move() {
	fmt.Println("zoumaobu")
}
