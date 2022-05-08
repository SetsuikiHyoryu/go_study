package main

import "sync"

// 声明互斥锁
var mutex sync.Mutex

func main() {
	mutex.Lock()
	defer mutex.Unlock()
}
