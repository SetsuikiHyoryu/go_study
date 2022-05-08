package main

import "fmt"

func main() {
	// 被 defer 的内容会在函数返回前执行
	defer func() {
		_panic := recover()

		if _panic != nil {
			fmt.Println(_panic)
		}
	}()

	panic("panic happened!")
}
