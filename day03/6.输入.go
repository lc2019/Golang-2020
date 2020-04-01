package main

import "fmt"

func main() {
	var (
		name string
		age  int
		marr bool
	)
	fmt.Println("请输入内容：")
	//fmt.Scan(&name,&age,&marr)
	fmt.Scanln(&name, &age, &marr)
	//fmt.Scanf("name:%s age: %d marr: %t\n",&name,&age,&marr)
	fmt.Println(name, age, marr)
}
