package main

import "fmt"

func main() {
	left := disc()
	fmt.Println(left)
}

var (
	coins = 50
	user  = []string{"Leic", "hello", "LovE", "lulu"}
	dis   = make(map[string]int, len(user))
)

func disc() int {
	for _, name := range user {
		for _, bt := range name {
			switch bt {
			case 'i', 'I':
				coins--
				dis[name]++
			case 'e', 'E':
				coins -= 2
				dis[name] += 2
			case 'l', 'L':
				coins -= 3
				dis[name] += 3
			case 'u', 'o':
				coins -= 4
				dis[name] += 4
			}
		}
	}
	for key, value := range dis {
		fmt.Println(key, value)
	}
	return coins
}
