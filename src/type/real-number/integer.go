package main

import "fmt"

func main() {
    // int 五种有负数的整数之一
	year := 2022
    // uint 五种没有负数的整数之一
	var month uint = 2
    // %T 可以打印数据类型
	fmt.Printf("%T 年 %T 月\n", year, month)
	fmt.Printf("%d 年 %d 月\n", year, month)
    
    fmt.Println("===")
    
    // 使用 uint8 表示颜色
    var red, green, blue uint8 = 0, 141, 213
    fmt.Println(red, green, blue)
    // 十六进制表示法
    var red16, green16, blue16 uint8 = 0x00, 0x8d, 0xd5
    // 用 0 填充，最小宽度为 2
    fmt.Printf("color: #%02x%02x%02x\n", red16, green16, blue16)
    
    fmt.Println("===")
    
    // 整数环绕（最大值时进位时超出内存，所以变为了 0）
    var number uint8 = 255
    number++
    fmt.Println(number)
    
    var number02 int8 = 127
    // %b 打印 bit
    fmt.Printf("%08b: ", number02)
    fmt.Println(number02)
    number02++
    fmt.Printf("%08b: ", number02)
    fmt.Println(number02)
}