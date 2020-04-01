package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//序列化
	j1 := js{"lc", 18}
	//func Marshal(v interface{}) ([]byte, error)
	//返回byte和err
	v, _ := json.Marshal(j1) //序列化的结果小写
	fmt.Println(string(v))

	//反序列化
	s := &js{}
	//func Unmarshal(data []byte, v interface{}) error {
	//接受byte切片
	json.Unmarshal([]byte(v), s)
	fmt.Println(s)
}

//元信息 json  tag 显示
type js struct {
	Name string `json:"name"` //必须首字母大写外部才能访问
	Age  int    `json:"age"`  //tag
}
