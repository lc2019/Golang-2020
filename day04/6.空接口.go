package main

import "fmt"

//空接口的类型可以是任意类型
func main() {
	var x interface{}
	x = 100
	fmt.Println(x)

	x = "lc"
	fmt.Println(x)

	x = false
	fmt.Println(x)
	showType(x)
	showType(100)
	showType("lc")

	//空接口的map
	stu := make(map[string]interface{}, 10)
	stu["lc"] = 100
	stu["lu"] = "2020"
	fmt.Println(stu)

	//类型断言
	pdt("lc")

}
func showType(a interface{}) {
	fmt.Printf("type :%T", a)
	fmt.Println()
}

func pdt(y interface{}) {
	//类型断言 空接口.(type)
	switch y.(type) {
	case int:
		fmt.Println("int", y)
	case string:
		fmt.Println("string", y)
	case bool:
		fmt.Println("bool", y)
	default:
		fmt.Println("cuowu")
	}
}
