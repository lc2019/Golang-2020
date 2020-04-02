package main

import (
	"fmt"
	"os"
)

func main() {
	//创建目录
	file := "./test_go"
	err := os.Mkdir(file, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}

	//创建目录
	file2 := "./test_go/demo/"
	err = os.MkdirAll(file2, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file2)
	}

	//创建文件,如果文件存在直接覆盖
	file3, err := os.Create("./test_go/demo.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(file3)
	}

	//打开文件
	file4, err := os.Open("./test_go/demo.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(file4.Name())
	}

	//打开文件openfile
	file5 := "./test_go/demo1.txt"
	f5, err := os.OpenFile(file5, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(f5.Name())
	}
	f5.Close()

	//删除文件remove，removeall
	err = os.Remove("./test_go/demo1.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ok")
	}

}
