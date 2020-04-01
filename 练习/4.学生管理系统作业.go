package main

import (
	"fmt"
	"os"
)

//学生管理系统作业

func main() {
	//学生管理者初始化
	smr := newStuMgr()
	for {
		sMenu()
		//用户输入
		var op int
		fmt.Println("请输入选项：")
		fmt.Scanln(&op)
		switch op {
		case 1:
			//用户输入信息
			stu := input()
			//学生管理者的方法
			smr.addStu(stu)
		case 2:
			stu := input()
			smr.modiStu(stu)
		case 3:
			stu := input()
			smr.delStu(stu)
		case 4:
			smr.showInfos()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("cuowu")
		}
	}
}

//0.用户输入构成结构体
func input() *stu {
	var (
		name string
		id   int64
		age  int
	)
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入学号")
	fmt.Scanln(&id)
	stu := newStu(name, age, id)
	return stu
}

//1.学生信息结构体
type stu struct {
	name string
	age  int
	id   int64
}

//2.学生信息构造
func newStu(name string, age int, id int64) *stu {
	return &stu{name: name, age: age, id: id}
}

//3.学生管理者的结构体
type stuMgr struct {
	allstu []*stu
}

//4.学生管理者的构造-初始化
func newStuMgr() *stuMgr {
	return &stuMgr{
		allstu: make([]*stu, 0), //初始化
	}
}

//5.菜单
func sMenu() {
	fmt.Println("学生信息管理系统")
	fmt.Println("	1.添加学生")
	fmt.Println("	2.修改学生")
	fmt.Println("	3.删除学生")
	fmt.Println("	4.显示学生")
	fmt.Println("	5.退出")
}

//6.添加学生
func (s *stuMgr) addStu(stu *stu) {
	for _, v := range s.allstu {
		if v.name == stu.name {
			fmt.Println("yi cunzai")
		}
	}
	//添加
	s.allstu = append(s.allstu, stu)
}

//7.修改学生
func (s *stuMgr) modiStu(stu *stu) {
	for key, value := range s.allstu {
		if value.name == stu.name {
			//修改
			s.allstu[key] = stu
			return
		}
	}
	fmt.Println("bu cunzai")
}

//8.删除学生
func (s *stuMgr) delStu(stu *stu) {
	for k, v := range s.allstu {
		if v.name == stu.name {
			s.allstu = append(s.allstu[:k], s.allstu[k+1:]...)
			return
		}
	}
	fmt.Println("bu cunzai ")
}

//9.显示学生
func (s *stuMgr) showInfos() {
	for _, v := range s.allstu {
		fmt.Println(v.name, v.age, v.id)
	}
}
