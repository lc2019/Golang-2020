package main

import "fmt"

func main() {
	var a NewI
	var b myInt
	a = 10 //main.newI
	b = 10 //int
	fmt.Printf("%T %T\n", a, b)

	s1 := stu{
		name: "lc",
		age:  18,
	}
	fmt.Println(s1, s1.age, s1.name)

}

// NewI 是一个int类型 注释
type NewI int
type myInt = int //类型别名，只在编写过程中

type stu struct {
	name string
	age  int
}
