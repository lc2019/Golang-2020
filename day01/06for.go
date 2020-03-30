package main

import "fmt"

func main() {
	//模式1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	s := "hello world"
	for i, v := range s {
		fmt.Printf("%d %c", i, v)
	}
}
