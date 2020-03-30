package main

import "fmt"

//字符串的修改转换
func main() {
	s := "hello world"
	n := len(s)

	for i := 0; i < n; i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
	for _, s := range s { //range 出来时是byte
		fmt.Printf("%c", s)
	}
	fmt.Println()

	//类型转换
	s1 := "雷超超"
	s3 := []rune(s1) //把字符串转换成切片
	s3[0] = '嗨'      //''表示字符 ""表示字符串
	fmt.Println(string(s3))

	//字符串和字节
	s5 := "h"
	s4 := 'g'
	fmt.Printf("%T %v\n", s5, s5) //string
	fmt.Printf("%T %c\n", s4, s4) //int32

	//字符串的翻转
	s1 = "hello"
	bs1 := []byte(s1) //h e l l o  前与后互换位置
	s2 := ""
	for i := len(bs1) - 1; i >= 0; i-- {
		s2 += string(bs1[i]) //字符串不能直接和unicode编码相加
		//s2 +=  //字符串不能直接和unicode编码相加
	}
	fmt.Println(s2)

	//方法2
	len1 := len(bs1)
	for i := 0; i < len1/2; i++ {
		bs1[i], bs1[len1-1-i] = bs1[len1-1-i], bs1[i]
	}
	fmt.Println(string(bs1))

}
