package main

import "fmt"

// 定义类型
type operator string

// 将方法关联到类型
func (loli operator) sleep() {
	fmt.Printf("%v is sleeping in my arms.\n", loli)
}

func main() {
	var loli operator = "Shamare"
	// 使用关联后的方法
	loli.sleep()
}
