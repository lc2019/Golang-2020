package main

import (
	"fmt"
)

func main() {
	res := splus("97", "2")
	fmt.Println(res)
}

func splus(s1, s2 string) (res string) {
	if len(s1) == 0 && len(s2) == 0 {
		res = "0"
		return
	}

	index1 := len(s1) - 1
	index2 := len(s2) - 1
	var left int //进位

	//97 3
	for index1 >= 0 && index2 >= 0 {
		c1 := s1[index1] - '0'
		c2 := s2[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		//遇到整除的情况 10 100%10 =0
		c3 := (sum % 10) + '0' //余数
		res = fmt.Sprintf("%c%s", c3, res)
		index1--
		index2--
	}

	//长的部分
	for index1 >= 0 {
		c1 := s1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' //余数
		res = fmt.Sprintf("%c%s", c3, res)
		index1--

	}
	for index2 >= 0 {
		c2 := s2[index2] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		c3 := (sum % 10) + '0' //余数
		res = fmt.Sprintf("%c%s", c3, res)
		index2--
	}

	if left == 1 { //拼接1+0
		res = fmt.Sprintf("1%s", res)
	}
	return
}
