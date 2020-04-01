package main

import "fmt"

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

//执行顺序：返回值赋值  defer   rerurn
func f1() int { //5
	x := 5
	defer func() { //没有返回值
		x++
	}()
	return x //返回值 = 5 x ++ ret
}
func f2() (x int) { //6
	defer func() {
		x++
	}()
	return 5 //x =5 x++ return x
}
func f3() (y int) { //5
	x := 5
	defer func() {
		x++
	}()
	return x //x = y = 5   return x
}
func f4() (x int) { //5
	defer func(x int) {
		x++
	}(x)
	return 5 //x=5 return
}
