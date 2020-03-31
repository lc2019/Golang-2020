package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	//map声明 map使用必须初始化
	var a map[string]int //nil

	//初始化
	a = make(map[string]int, 20)
	fmt.Println(a)

	a["leichao"] = 18
	fmt.Println(a)

	//声明同时初始化
	b := map[string]int{"lc": 18}
	fmt.Println(b)

	//判断某个键是否存在
	var stu = make(map[string]int, 8)
	stu["lc"] = 60
	stu["ll"] = 80

	//判断是否存在key
	v, ok := stu["lzc"] //如果存在v返回value，ok返回true 否则返回0 和false
	if ok {
		fmt.Println("ll ok", v, ok)
	} else {
		fmt.Println("bu cunz", v, ok)
	}

	//遍历map
	for k, v1 := range stu {
		fmt.Println(k, v1)
	}

	//添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%2d", i) //格式化打印
		value := rand.Intn(100)         //100以内的随机数
		stu[key] = value
	}

	//按照key进行排序
	s := []string{} //声明空的切片用来存储key
	for k := range stu {
		s = append(s, k) //保存到切片中
	}
	//fmt.Println(s)
	//对切片进行排序
	sort.Strings(s)
	for _, v2 := range s {
		fmt.Println(v2, stu[v2])
	}

	//类型为map的切片
	var ms = make([]map[string]int, 5) //切片初始化
	//需要给map进行初始化
	ms[0] = make(map[string]int, 8)
	ms[0]["lc"] = 100 //第一个切片的键值对
	fmt.Println(ms)   //[map[lc:100] map[]]

	//切片的map
	var sm = make(map[string][]int, 3) //map初始化
	//切片初始化
	sm["zg"] = make([]int, 8)
	sm["zg"][0] = 100 //值的第一个元素
	sm["zg"][1] = 200
	fmt.Println(sm) //[zg:[100 200 0]]

	//统计
	var s1 = "how do you do"
	count := map[string]int{}
	words := strings.Split(s1, " ") //转换成切片
	for _, w := range words {
		v, ok := count[w]
		if ok {
			count[w] += v //如果有则增加
		} else {
			count[w] = 1 //如果count中没有则赋值为1
		}
	}
	fmt.Println(count)
}
