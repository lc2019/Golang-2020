package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	f, err := os.Stat("/Users/mac/下载地址.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(f.Name(), f.Size(), f.IsDir(), f.Mode())
	}

	//路径相关
	fname := "/Users/mac/下载地址.txt"
	fmt.Println(filepath.IsAbs(fname)) //绝对路径

	fmt.Println(filepath.Abs(fname)) //绝对路径
	//路径拼接
	fmt.Println(path.Join("/Users/mac", fname))

}
