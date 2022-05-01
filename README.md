# Golang

## 教程地址

[Go语言编程快速入门（Golang）（完结）](https://www.bilibili.com/video/BV1fD4y1m7TD)

## 环境搭建

### 官网

- 境外：<https://go.dev/>
- 境内：<https://golang.google.cn>

※ 注意需要配置 GOPATH

### VSCode

1. 安装 `golang.go` 插件
2. 执行 `>go install/update Tools` 命令
   - 安装目录为 `GOPATH`

#### 定制对 go 后缀文件的操作

```json
"[go]": {
  // 不将制表符转换为空格
  "editor.insertSpaces": false,
  // 制表符所占大小为 4 个空格的宽度
  "editor.tabSize": 4,
  // 设置默认的格式化工具为 goloang.go
  "editor.defaultFormatter": "golang.go",
}
```

## 创建 go 文件

|名称|内容|
|-|-|
|后缀|go|
|编译并运行|go run \<filename\>|

## package & function

```go
// main.go

// 声明代码所在的包
package main

// 指明代码所需的包
import (
  "fmt" // 包中包含很多函数
)

// main 函数是 go 语言的入口函数，如果不定义就会报错
// 左大括号必须和 func 在同一行，否则报错
func main() {
  fmt.Println("Hello World!")
}
```

## 格式化打印

```go
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


```

```powershell
$ go run main.go
shamare's height is 138cm
shamare    $ 138cm
sora       $ 155cm
```

## 声明变量

|名称|关键字|说明|
|-|-|-|
|常量|const|值不可改变|
|变量|var|值可改变|
|变量短声明|:=|不可以用来定义全局变量|

### 复数声明

```go
package main

import "fmt"


func main()  {
    // 复数声明的方式
    var(
        shamare = "shamare"
        shamareHeight = 138
    )
    
    // 方式二
    var suzuran, suzuranHeight = "suzuran", 137
    
    fmt.Println(shamare, shamareHeight)
    fmt.Println(suzuran, suzuranHeight)
}

```

## 循环与分支

- 条件不需要括号
- 空字符串不可作为判断条件
- `switch` 和 `if` 的条件所用到的变量可以在判断语句中声明
  - 仅可使用 `:=` 的形式

### strings.Contains

判断参数一的字符串中是否包含参数二的字符串

### switch

- `switch` 的 `case` 可以填复数值，用 `,` 分隔
- `case` 不用相对于 `switch` 缩进

```go
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
```

```powershell
$ go run switch.go
shamare is a fox.
suzuran is a fox.
```

### for

- for 之后不跟条件即为无限循环
- 可以用 `break` 手动跳出循环
- 只写一个式子默认为是第二区域的式子

## 类型

### 实数

#### 浮点数

- 浮点型默认为 `float64`
- 32 为单精度（4 字节）；64 为双精度（8 字节）
  - 处理诸如 3D 游戏中的大量数据时，可以牺牲精度节省内存

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // 浮点型格式化动词
    floatSample := 1.0 / 3
    
    fmt.Printf("v: %v\n", floatSample)
    fmt.Printf("f: %f\n", floatSample)
    fmt.Printf(".3f: %.3f\n", floatSample)

    // 宽度.精度。此处意为宽度 4，精度为保留 2 位
    // 宽度大于数字长度时，左边填充空格；不指定宽度时，按实际长度显示
    // 如果要填充零，可以写作 `%04.2f`
    fmt.Printf("4.2f: %4.2f\n", floatSample)
    
    // 比较浮点型
    piggyBank := 0.1
    piggyBank += 0.2
    fmt.Println(piggyBank == 0.3) // false
    // 正确比较方式
    fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
}

```

```powershell
$ go run realNumber.go 
v: 0.3333333333333333
f: 0.333333
.3f: 0.333
4.2f: 0.33
false
true
```

#### 整数

##### 十种整数类型（本质是八种）

|Type|Range|Storage|
|-|-|-|
|int8|-128 - 127|8-bit(one byte)|
|uint8|0 - 255|8-bit(one byte)|
|int16|-32,768 - 32767|16-bit(two bytes)|
|uint16|0 - 65535|16-bit(two bytes)|
|int32|-2,147,483,648 - 2147,483,647|32-bit(four bytes)|
|uint32|0 - 4,294,967,295|32-bit(four bytes)|
|int64|-9,223,372,036,854,775,808 - 9,223,372,036,854,775,807|64-bit(eight bytes)|
|uint64|0 - 18,446,644,073,709,551,615|64-bit(eight bytes)|

※ int, uint 在老的移动设备上是 32 位，在新的计算机上是 64 位。

```go
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

```

```powershell
$ go run integer.go
int 年 uint 月
2022 年 2 月
===
0 141 213
color: #008dd5
===
0
01111111: 127
-10000000: -128
```

##### big

- 使用 big.Int 进行运算时，等式中的其他部分都必须是 big.Int
- `NewInt()` 函数可以把 `int64` 转化为 `big.Int` 类型

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 指数的值如果不主动声明类型则类型为 float64
	var distance = 24e18
	fmt.Println(distance)

	// 使用 big 包
	distance02 := new(big.Int)
	distance02.SetString("24000000000000000000", 10)
	fmt.Println(distance02)

	// 常量可以无类型（untype）
	const distance03 = 24000000000000000000000000
	// 程序中的每个字面值都是常量
	// 针对字面值的常量的计算在编译阶段即完成
	fmt.Println(24000000000000000000000000 / 299792 / 86400)

}

```

```powershell
$ go run bigPackage.go 
2.4e+19
24000000000000000000
926568346646267
```

### 多语言文本

- 字符串字面值：`"\n"`，即可以包含转义。
- 原始字符串字面值：\`\n\`

|名称|含义|
|-|-|
|字符|比如 65 代表 A 字符|
|code point|Unicode 联盟为超过 100 万个字符分配了相应的数值，这个数叫作 code point|
|rune|为了表示 unicode code point，go 语言提供的 int32 的类型别名|
|byte|uint8 的类型别名，目的是用于二进制数|

#### 类型别名

`type byte = uint8`

```go
package main

import "fmt"

func main() {
	// 任何整数类型都可以使用 `%c` 打印字符，但是 `rune` 可以指明该数值是用来表示一个字符
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	// 打印出的是 code point
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	// 打印出的是字符
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
}

```

```powershell
$ go run main.go
960 940 969 33
π ά ω !
```

#### 字符

```go
package main

import "fmt"

func main() {
	// 字符字面值使用 `''` 包裏
	// 不指定类型的话会被推断为 rune
	grade := 'A'
	fmt.Printf("grade: %T %v %c", grade, grade, grade)
}
```

```powershell
$ go run rune.go
grade: int32 65 A
```

#### string

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	operator := "巫恋"

	// len 输出字节数而不是字符数
	fmt.Println(len(operator), "bytes")
	// utf8 的包带有输出字符数的方法
	fmt.Println(utf8.RuneCountInString(operator), "runes")

	// DecodeRundeInString 返回字符串中的第一个字符和它的字节数
	runeItem, size := utf8.DecodeRuneInString(operator)
	fmt.Printf("first rune: %c, %v bytes\n", runeItem, size)

	println("===")

	// range 返回的 index 是以根据字节数计算的
	for index, item := range operator {
		fmt.Println(index, item)
		fmt.Printf("%v, %c\n", index, item)
		fmt.Println("===")
	}
}

```

```powershell
$ go run string.go
6 bytes
2 runes
first rune: 巫, 3 bytes
===
0 24043
0, 巫
===
3 24651
3, 恋
===
```

### 类形转换

- 无符号与有符号的整数型之间也需要转换。
- 不同大小的整数型之间也需要转换。
- 将整数型转化为字符串型时，如果它的值不能转化为 code point，虽然不会报错，但会輸出乱码。
- 整数型与布尔型，字符串型与布尔型之间无法互相转换

```go
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	description := "Shamare is" + "138cm"
	fmt.Println(description)
	// 直接连接字符串与数值会报错
	// description02 := "Shamare is " + 138 + "cm"

	// 整数型与浮点型也不能混合使用
	pi := math.Pi
	times := 2
	// 报错
	// fmt.Println(pi * times)
	fmt.Println(pi, times)

	// 类形转换使用目标类型去包裏
	// 从浮点型转换为整数型时，小数点之后的部分会被截断而不是舍入
	fmt.Println(int(pi))

	// rune, bype 可以转换成 string，输出结果为该值对应的字符
	var piItme rune = 960
	var bang byte = 33
	fmt.Println(string(piItme), string(bang))
	// 如果需要直接输出字符串型的数字，需要使用 strconv 包的 Itoa 函数
	// Itoa: Integer to ASCII
	description02 := "Shamare is " + strconv.Itoa(138) + "cm"
	fmt.Println(description02)

	// strconv 包中亦有 Atoi 函数，但需要错误处理
	text, exception := strconv.Atoi("10")

	if exception != nil {
		fmt.Println(exception.Error())
		return
	}

	fmt.Println(text)
}

```

```powershell
$ go run typeChange.go 
Shamare is138cm
3.141592653589793 2
3
π !
Shamare is 138cm
10
```

### 声明新类型

- 可以使用 `type` 关键字来声明类型。
- 自定义类型可以提高代码可读性。

## 函数

- 大写字母开头的函数、变量或其他标识符被被导出。

```go
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

```

```powershell
$ go run parameterChangeableFunction.go 
Shamare is sleeping in my arms.
Shamare, Suzuran are sleeping in my arms.
```

### 方法

- 函数是独立存在的。
- 方法是指与类型相关联的函数。

```go
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

```

```powershell
$ go run method.go 
Shamare is sleeping in my arms.
```

### 声明函数类型

```go
package main

import "fmt"

func main() {
	var sleep sleep = func(name string) {
		fmt.Printf("%v is sleeping in my arms.\n", name)
	}

	sleep("Shamare")
}

// 声明函数类型
type sleep func(name string)

```

```powershell
$ go run functionType.go 
Shamare is sleeping in my arms.
```

### 闭包

```go
package main

import "fmt"

func main() {
	count := 0

	visitCount := func() int {
		return count
	}

	fmt.Println(visitCount()) // 0

	count++
	fmt.Println(visitCount()) // 1
}

```

```powershell
$ go run closure.go 
0
1
```
