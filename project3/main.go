package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 使用带缓冲的通道（容量1）避免阻塞
	oddSignal := make(chan struct{}, 1)
	evenSignal := make(chan struct{}, 1)

	// 奇数协程（打印1,3,5,7,9）
	go func() {
		defer wg.Done()
		for num := 1; num <= 10; num += 2 {
			<-oddSignal // 等待奇数打印信号（缓冲通道不会阻塞）
			fmt.Printf("奇数: %d\n", num)
			evenSignal <- struct{}{} // 发送偶数打印信号（缓冲通道不会阻塞）
		}
	}()

	// 偶数协程（打印2,4,6,8,10）
	go func() {
		defer wg.Done()
		for num := 2; num <= 10; num += 2 {
			<-evenSignal // 等待偶数打印信号（缓冲通道不会阻塞）
			fmt.Printf("偶数: %d\n", num)
			oddSignal <- struct{}{} // 发送奇数打印信号（缓冲通道不会阻塞）
		}
	}()

	// 启动初始信号（触发第一个奇数打印）
	oddSignal <- struct{}{}

	// 等待所有协程完成
	wg.Wait()
}
