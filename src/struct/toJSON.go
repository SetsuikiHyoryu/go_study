package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// 要转换成 JSON 时属性名必须大写，否则属性不能导出
	type operator struct {
		// 自定义 JSON 输出时的格式
		Name   string `json:"name"`
		Height int    `json:"height"`
	}

	shamare := operator{Name: "Shamare", Height: 138}

	bytes, exception := json.Marshal(shamare)
	handleError(exception)

	fmt.Println(shamare)
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}

func handleError(exception error) {
	if exception == nil {
		return
	}

	fmt.Println(exception)
	// 退出程序
	os.Exit(1)
}
