package main

import "fmt"

func main() {
	str := "_  asvc123!"
	w, s, n, o := count(str)
	fmt.Println(w, s, n, o)
}

func count(s string) (word, space, num, other int) {
	str := []rune(s)
	for _, v := range str {
		switch {
		//大小写字母一起统计
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			word++
		case v == ' ':
			space++
		case v >= '0' && v <= '9':
			num++
		default:
			other++
		}
	}
	return
}
