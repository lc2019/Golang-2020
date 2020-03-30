package main

import "fmt"

func main() {
	res := spl("123", "47")
	fmt.Println(res)
}

func spl(s1, s2 string) (res string) {
	//如果2数长度为0结束
	if len(s1) == 0 && len(s2) == 0 {
		res = "0"
		return
	}
	//下标
	index1 := len(s1) - 1
	index2 := len(s2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {
		c1 := s1[index1] - '0'
		c2 := s2[index2] - '0'

		//和
		sum := int(c1) + int(c2) + left
		//进位
		c3 := (sum % 10) + '0'
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		res = fmt.Sprintf("%c%s", c3, res)
		index2--
		index1--

	}

	for index1 >= 0 {
		c1 := s1[index1] - '0'
		sum := int(c1) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
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
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)

		index2--
	}

	if left == 1 {
		res = fmt.Sprintf("1%s", res)
	}

	return
}
