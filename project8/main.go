package main

import (
	"fmt"
	"sync"
)

// 生产者协程：向带缓冲通道发送 1-100 的整数
// ch: 带缓冲的整数通道（发送端）
// wg: 用于同步协程完成的等待组
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后通知 WaitGroup

	// 生成 1 到 100 的整数并发送到通道
	for i := 1; i <= 100; i++ {
		ch <- i // 发送到带缓冲通道（缓冲区未满时不会阻塞）
		fmt.Printf("生产者发送: %d（缓冲区剩余容量: %d）\n", i, cap(ch)-len(ch))
	}

	// 发送完所有数据后关闭通道（必须由发送方关闭）
	close(ch)
	fmt.Println("生产者完成所有数据发送，通道已关闭")
}

// 消费者协程：从带缓冲通道接收整数并打印
// ch: 带缓冲的整数通道（接收端）
// wg: 用于同步协程完成的等待组
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后通知 WaitGroup

	// 循环接收通道数据，直到通道关闭（ok 为 false 时退出）
	for num := range ch {
		fmt.Printf("消费者接收: %d\n", num)
	}
	fmt.Println("消费者检测到通道关闭，停止接收")
}

func main() {
	// 创建带缓冲的通道（容量为 10）
	// 缓冲区大小可根据需求调整（此处设为 10）
	bufferSize := 10
	ch := make(chan int, bufferSize)

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
