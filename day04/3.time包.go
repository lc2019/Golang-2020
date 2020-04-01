package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //time的结构体
	fmt.Println(now.Year())
	fmt.Println(now.Unix()) //时间戳
	//时间格式是go语言的诞生日期
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))

	tform(now)

}
func tform(t time.Time) {
	timestr := t.Format("2006-01-02 15:04:05")
	fmt.Println(timestr)
	calcTime()
}

func calcTime() {
	start := time.Now().UnixNano() / 1000
	s2 := time.Now()
	time.Sleep(300)
	fmt.Println("xiuxi shijian")
	end := time.Now().UnixNano() / 1000
	//内置的函数
	fmt.Println("耗费时间：", time.Since(s2)) //现在时间减去给的时间
	fmt.Println("耗费时间：", end-start)

}
