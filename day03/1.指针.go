package main

import "fmt"

func main() {
	//具体类型的地址
	var a int
	b := &a //取地址
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)

	d := 100
	b = &d
	fmt.Println(*b)

	//*根据地址取值 &取地址

	a1 := [3]int{1, 2, 3}
	modif(a1)
	modif2(&a1)
	fmt.Println(a1)

}

//指针的应用
func modif(a [3]int) {
	a[0] = 100
}
func modif2(a *[3]int) {
	a[0] = 100
}
