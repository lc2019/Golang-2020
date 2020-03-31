package main

import "fmt"

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

}
