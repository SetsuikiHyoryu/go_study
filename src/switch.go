package main

import "fmt"

func main() {
    var operator = "shamare"
    
    switch operator {
    case "sora":
        fmt.Println("sora is a bunny.")
    case "shamare":
        fmt.Println("shamare is a fox.")
        // 使用此关键字可执行下一 case 中的代码
        fallthrough
    case "suzuran":
        fmt.Println("suzuran is a fox.")
    case "rope":
        fmt.Println("rope is a bunny.")
    }
}