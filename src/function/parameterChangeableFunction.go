package main

import "fmt"

func main() {
	sleep("Shamare")
	sleep("Shamare", "Suzuran")
}

// `...` 表示函数的参数可变；`interface{}` 为空接口，所有类型均可实现
func sleep(nameList ...interface{}) {
	length := len(nameList)
	var name string

	for i := 0; i < length; i++ {
		name += nameList[i].(string)

		if i != length-1 {
			name += ", "
		}
	}

	beVerb := "is"

	if length > 1 {
		beVerb = "are"
	}

	fmt.Printf("%v %v sleeping in my arms.\n", name, beVerb)
}
