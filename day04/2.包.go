package main

/*
放在gopath src目录的文件夹pkg下，函数名大写
*/

import (
	"fmt"
	"pkg"
)

func main() {
	fmt.Println(pkg.Add(100, 200))
}
