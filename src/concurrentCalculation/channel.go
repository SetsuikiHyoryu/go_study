package main

import (
	"fmt"
	"time"
)

func main() {
	channelShamare := make(chan string)
	channelSuzuran := make(chan string)
	var count int

	go shamareAttack(channelShamare)
	go suzuranAttack(channelSuzuran)

	for {
		// 使用 select 执行未被阻塞的通道
		select {
		case text, opening := <-channelShamare:
			if !opening {
				return
			}

			fmt.Println(text)
			count++

		case text, opening := <-channelSuzuran:
			if !opening {
				return
			}

			fmt.Println(text)
			count++
		}

		fmt.Printf("Attaced %v times.\n", count)
	}

}

func shamareAttack(channel chan string) {
	text := "Shamare attacked."

	for i := 0; i < 3; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	// 使用 close 关闭通道以避免阻塞死锁
	close(channel)
}

func suzuranAttack(channel chan string) {
	text := "Suzuran attacked."

	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Millisecond * 500)
	}

	close(channel)
}
