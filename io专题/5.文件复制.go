package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcfile := "./test_go/demo.txt"
	desfile := "./test_go/demo2.txt"
	//n, err := copyfile2(srcfile, desfile)
	n, err := copyfile(srcfile, desfile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
func copyfile(srcfile, desfile string) (int, error) {
	file1, err := os.Open(srcfile)
	if err != nil {
		return 0, err
	}
	//注意创建文件的权限
	file2, err := os.OpenFile(desfile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	//读取的字节
	bs := make([]byte, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if n == 0 && err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		} else {
			total += n
			file2.Write(bs[:n])
		}
	}
	return total, nil
}
func copyfile2(srcfile, desfile string) (int64, error) {
	file1, err := os.Open(srcfile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(desfile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}
