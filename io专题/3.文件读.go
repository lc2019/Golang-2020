package main

import (
	"fmt"
	"io"
	"os"
)

//读取文件，复制文件 []byte
//读取文件结束 返回 n err n =0 err =io.eof

//写入文件 write（[]byte） writestring（）byte
func main() {
	//打开文件
	filename := "./test_go/demo.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		sli := make([]byte, 1024) //每次读取的数量切片
		//循环读取
		n := -1
		for {
			n, err = f.Read(sli)
			//读取结束 返回0和io.eof
			if n == 0 && err == io.EOF {
				fmt.Println("读取结束")
				break //循环退出
			}
			//每次读取的内容截取
			fmt.Println(string(sli[:n]))
		}
	}
	f.Close() //流关闭
}
