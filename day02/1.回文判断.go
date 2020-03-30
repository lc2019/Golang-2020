package main

import "fmt"

func main() {
	str := "abc上海自来水来自海上cba"
	if pro(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func pro(str string) bool {
	t := []rune(str)
	for i, _ := range t {
		//判断1半 因为回文对称
		if i == len(t)/2 {
			break
		}

		last := len(t) - 1 - i
		//判断是否相等
		if t[i] != t[last] {
			return false
		}
	}
	return true
}
