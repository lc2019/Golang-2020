package main

import "fmt"

func main() {
	d := dog{1, &animal{"wangwang"}}
	d.move()
	fmt.Println(d.name, d.feat)
}

type animal struct {
	name string
}

type dog struct {
	feat    int
	*animal //模拟继承animal的字段-name
}

func (d *dog) move() {
	fmt.Println("paoqilai", d.name)
}
