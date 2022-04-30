package main

import "fmt"

func main() {
    // 使用 `fmt.Printf`，第一个参数必须是字符串。
    // 接受如 `%v` 的格式化动词，它的值将由第二个参数替换。
    // 传入复数格式化动词时，值将由后续参数按顺序替换。
    fmt.Printf("%v's height is %vcm\n", "shamare", 138)
    
    // 前置数字可以向左右填充空格以自动对齐，负右正左。
    // 文字超出指定数填充数时将不填充一部分。
    fmt.Printf("%-10v $%4vcm\n", "shamare", 138)
    fmt.Printf("%-10v $%4vcm\n", "sora", 155)
}
