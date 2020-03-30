package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(pld(i))
	}
}
func pld(n int) int {
	if n <= 1 {
		return 1
	}
	return pld(n-1) + pld(n-2)
}
