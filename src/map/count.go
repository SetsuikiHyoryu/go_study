package main

import "fmt"

func main() {
	supporter := "Supporter"
	medic := "Medic"

	operators := map[string]string{
		"Shamare":    supporter,
		"Suzuran":    supporter,
		"Honeyberry": medic,
	}

	// 创建一个计数器
	classCount := make(map[string]int)

	for _, class := range operators {
		// 遇到相同的 class 就会叠加
		classCount[class]++
	}

	// 遍历计数器
	for class, number := range classCount {
		fmt.Printf("%v: %v\n", class, number)
	}
}
