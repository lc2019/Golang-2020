package main

import (
	"fmt"
	"os"
)

/*
需求分析：任务分解 伪代码
0.定义结构体
1.打印菜单
2.等待用户输入
3，添加书籍
4.修改书籍
5.显示书籍
6.退出
*/
func main() {
	for {
		showMenu()
		//2.等待用户输入
		var op int
		fmt.Scanln(&op)
		switch op {
		case 1:
			addBook()
		case 2:
			updateBook()
		case 3:
			showBook()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("shuru cuowu")

		}
	}

}

//0.定义书籍结构体
type book struct {
	title    string
	author   string
	price    int
	publishd bool
}

//1.所有书籍的切片--指针类型的切片
var allBooks = make([]*book, 0, 200)

//2.创建新书的构造函数
func newBook(title, author string, price int, publishd bool) *book {
	return &book{
		title:    title,
		author:   author,
		price:    price,
		publishd: publishd,
	}
}

//3.打印菜单
func showMenu() {
	fmt.Println("欢迎来到书籍管理系统")
	fmt.Println("	1.添加书籍")
	fmt.Println("	2.修改书籍")
	fmt.Println("	3.显示书籍")
	fmt.Println("	4.退出")
}

//4.获取用户输入-返回书籍的指针
func userInput() *book {
	//3.1获取用户输入
	var (
		title    string
		author   string
		price    int
		publishd bool
	)
	fmt.Println("请根据提示输入内容：")
	fmt.Println("请输入书名：")
	fmt.Scanln(&title)
	fmt.Println("请输入作者：")
	fmt.Scanln(&author)
	fmt.Println("请输入价格：")
	fmt.Scanln(&price)
	fmt.Println("是否上架：")
	fmt.Scanln(&publishd)
	fmt.Println(title, author, price, publishd)
	book := newBook(title, author, price, publishd)
	return book
}

//5，添加书籍
func addBook() {
	book := userInput()
	for _, v := range allBooks {
		//指针类型的切片
		if v.title == book.title {
			fmt.Println("cunz", v.title)
			return
		}
	}
	allBooks = append(allBooks, book)
	fmt.Println("添加书籍成功")
}

//6.修改书籍
func updateBook() {
	book := userInput()
	//遍历所有的书籍
	for i, v := range allBooks {
		//如果书名相等
		if book.title == v.title {
			//allbooks是指针类型的切片
			allBooks[i] = book //等于输入的内容 --更新
			fmt.Println("更新ok")
			return
		}
	}
	fmt.Println("书名不存在")
}

//7.显示书籍
func showBook() {
	if len(allBooks) == 0 {
		fmt.Println("还没有书籍")
	}
	for _, v := range allBooks {
		//指针类型的切片
		fmt.Printf("%s %s %d %T\n", v.title, v.author, v.price, v.publishd)
	}
}
