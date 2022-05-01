package main

import "fmt"

func main() {
	var sleep sleep = func(name string) {
		fmt.Printf("%v is sleeping in my arms.\n", name)
	}

	sleep("Shamare")
}

// 声明函数类型
type sleep func(name string)
