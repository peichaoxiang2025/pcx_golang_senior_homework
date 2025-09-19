package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter int64          // 原子计数器（使用 int64 类型，支持 atomic 包操作）
		wg      sync.WaitGroup // 等待组（协调协程完成）
	)

	// 设置需要等待的协程数量（10 个）
	wg.Add(10)

	// 启动 10 个协程并发递增计数器
	for i := 0; i < 10; i++ {
		// 使用闭包避免循环变量捕获问题（确保每个协程独立执行 1000 次）
		go func() {
			defer wg.Done() // 协程完成后通知 WaitGroup

			// 每个协程对计数器执行 1000 次原子递增操作
			for j := 0; j < 1000; j++ {
				// 原子递增操作（等价于 counter++，但由 CPU 保证原子性）
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// 主协程等待所有子协程完成（阻塞直到 wg 计数器归零）
	wg.Wait()

	// 输出最终计数器值（预期为 10 * 1000 = 10000）
	fmt.Println("Final counter value:", counter)
}
