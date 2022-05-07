package main

import "fmt"

func main() {
	shamare := operator{name: "Shamare"}
	// 因为构造体中的字段实现了 operatorer 接口，所以构造体的实例也就实现了接口
	printName(shamare)
}

type operatorer interface {
	getName() name
}

// 需要传入实现了接口的参数
func printName(unit operatorer) {
	fmt.Println(unit.getName())
}

type operator struct {
	// 使用嵌入的方式定义构造体
	name
}

type name string

func (_name name) getName() name {
	return _name
}
