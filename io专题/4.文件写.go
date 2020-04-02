package main

import (
	"fmt"
	"os"
)

func main() {
	//1.打开文件
	file, err := os.OpenFile("./test_go/demo1.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		//写入的长度 n
		n, err := file.Write([]byte("1233"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("写入长度", n)
		}
		fmt.Println(n)
	}
}
