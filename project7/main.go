package main

import (
	"fmt"
	"sync"
)

// producer 协程：生成 1-10 的整数并发送到通道
// ch: 用于传递整数的通道（仅发送）
// wg: 用于同步协程完成的等待组
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后通知 WaitGroup

	// 生成 1 到 10 的整数并发送到通道
	for i := 1; i <= 10; i++ {
		ch <- i                      // 发送整数到通道（无缓冲通道会阻塞直到被接收）
		fmt.Printf("生产者发送: %d\n", i) // 可选：打印发送日志
	}

	// 发送完所有数据后关闭通道（关键操作！）
	// 避免消费者协程无限阻塞
	close(ch)
}

// consumer 协程：从通道接收整数并打印
// ch: 用于接收整数的通道（仅接收）
// wg: 用于同步协程完成的等待组
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后通知 WaitGroup

	// 循环接收通道数据，直到通道关闭（ok 为 false 时退出循环）
	for num := range ch {
		fmt.Printf("消费者接收: %d\n", num) // 打印接收到的整数
	}
}

func main() {
	// 创建无缓冲通道（用于协程间通信）
	// 无缓冲通道的特性：发送和接收操作必须同时准备好，否则阻塞
	ch := make(chan int)

	// 创建 WaitGroup 用于等待两个协程完成
	var wg sync.WaitGroup

	// 注册需要等待的协程数量（2 个）
	wg.Add(2)

	// 启动生产者协程（传递通道和 WaitGroup 指针）
	go producer(ch, &wg)

	// 启动消费者协程（传递通道和 WaitGroup 指针）
	go consumer(ch, &wg)

	// 主协程等待所有子协程完成（阻塞直到 wg 计数器归零）
	wg.Wait()

	// 主协程退出前打印提示
	fmt.Println("所有协程执行完成，程序退出")
}
