package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	var b = []int{1, 2, 3} //[]int
	fmt.Printf("%T%T%v%v", a, b, a, b)

	//切片引用类型
	c := b
	c[0] = 100
	fmt.Println(b, c)

	//切片的操作
	a1 := [...]int{1, 3, 5, 7, 8, 20}
	s1 := a1[1:4]
	fmt.Println(s1, len(s1), cap(s1))

	var a2 []int
	fmt.Printf("%d %d %p\n", len(a2), cap(a2), &a2)
	a2 = append(a2, 1, 2)
	fmt.Printf("%d %d %p\n", len(a2), cap(a2), &a2)
	a2 = append(a2, 1, 2, 3, 4)
	fmt.Printf("%d %d %p\n", len(a2), cap(a2), &a2)
	a2 = append(a2, 1)
	fmt.Printf("%d %d %p\n", len(a2), cap(a2), &a2)

	//copy 值拷贝 浅拷贝
	var s2 = []int{0, 0, 0}
	copy(s2, a2)
	s2[0] = 100
	fmt.Println(s2, a2)

	//切片元素的删除
	a3 := append(a2[:2], a2[3:]...)
	fmt.Println(a3)

	//以数组为切片的底层数组
	ss := a1[2:5]
	fmt.Println(ss)

}
