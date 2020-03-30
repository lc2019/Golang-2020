package main

import "fmt"

//字符串的修改转换
func main() {
	s := "hello world"
	n := len(s)

	for i := 0; i < n; i++ {
		//fmt.Println(s[i])
		fmt.Printf("%c\n", s[i])
	}

	for _, s := range s {
		fmt.Printf("%c\n", s)
	}

	//修改字符串
	s1 := "雷超超"
	s3 := []rune(s1) //把字符串转换成切片
	s3[0] = '嗨'      //''表示字符 ""表示字符串
	fmt.Println(string(s3))

	s5 := "h"
	s4 := 'g'
	fmt.Printf("%T %v\n", s5, s5) //string
	fmt.Printf("%T %c\n", s4, s4) //int32

}
