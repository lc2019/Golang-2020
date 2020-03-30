package main

import "fmt"

func main() {
	//模式1
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	s := "hello world"
	//模式2
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	//switch
	num := 3
	switch num {
	case 1:
		fmt.Println("1")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("qita")

	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto L1 //直接去执行l1下面的语句
			}
			fmt.Println(i, j)
		}
	}
L1:
	fmt.Println("跳出来了")
}
