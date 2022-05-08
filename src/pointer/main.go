package main

import "fmt"

func main() {
	operator := "Shamare"
	name := &operator
	fmt.Println(name)
	fmt.Println(*name)

	// 解引用也可以直接更改所引用的地址中的值
	*name = "Suzuran"
	fmt.Println(operator)

	type class struct {
		typeName string
	}

	// `&` 可以直接用在复合字面量之前
	_class := &class{typeName: "Supporter"}
	// 访问字段和数组时，解引用不是必须的
	_class.typeName = "medic"
	fmt.Printf("%+v\n", _class)
}
