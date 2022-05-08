package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	exception := writeFile("shamare")

	if exception != nil {
		fmt.Println(exception)
		os.Exit(1)
	}
}

type safeWriter struct {
	write     io.Writer
	exception error
}

func (_safeWriter safeWriter) writeln(text string) {
	if _safeWriter.exception != nil {
		return
	}

	// 每次错误发生时更新结构体中的错误
	_, _safeWriter.exception = fmt.Fprintln(_safeWriter.write, text)
}

func writeFile(fileName string) error {
	file, exception := os.Create(fileName)
	// 使用 defer 关键字，可以确保 deferred 的动作在函数返回前执行。
	// 即不用每个返回文前都添加相同的操作。
	defer file.Close()

	if exception != nil {
		return exception
	}

	_safeWriter := safeWriter{write: file}
	_safeWriter.writeln("Shamare is a supporter.")
	_safeWriter.writeln("Suzuran is a supporter too.")

	return _safeWriter.exception
}
