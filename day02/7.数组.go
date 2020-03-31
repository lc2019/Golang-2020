package main

import "fmt"

func main() {
	//数组声明
	var arr [3]int          //定义一个长度为3的int类型的数组
	fmt.Printf("%T\n", arr) //[3]int 数组长度是类型的一部分

	a1 := [3]int{0, 1, 2}
	fmt.Println(a1)

	//自动推导长度
	a2 := [...]int{1, 2, 3, 4}
	fmt.Println(a2)

	//数组求和
	a3 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a3 {
		sum += v
	}
	fmt.Println(sum)

	//找出数组中和为指定值的2个元素的下标
	for i, v := range a3 {
		other := 8 - v //和的另一半
		//从i后面开始计算和
		for i2 := i + 1; i2 < len(a3); i2++ {
			if a3[i2] == other {
				fmt.Println(i, i2)
			}
		}
	}

	//数组的类型
	a4 := a3
	a4[0] = 100
	fmt.Println(a3, a4)
}
