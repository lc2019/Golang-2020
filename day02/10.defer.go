package main

import "fmt"

func main() {
	de()
}

func de() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("re")
		}
	}()
	panic("haha")
}
