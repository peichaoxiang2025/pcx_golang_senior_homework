package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int            // 共享计数器（需要保护的共享资源）
		mu      sync.Mutex     // 互斥锁（保护计数器的原子性）
		wg      sync.WaitGroup // 等待组（协调协程完成）
	)

	// 设置需要等待的协程数量（10 个）
	wg.Add(10)

	// 启动 10 个协程并发递增计数器
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() // 协程完成后通知 WaitGroup

			// 每个协程对计数器执行 1000 次递增操作
			for j := 0; j < 1000; j++ {
				mu.Lock()   // 获取互斥锁（保护临界区）
				counter++   // 递增计数器（原子操作）
				mu.Unlock() // 释放互斥锁
			}
		}()
	}

	// 主协程等待所有子协程完成（阻塞直到计数器归零）
	wg.Wait()

	// 输出最终计数器值（预期为 10 * 1000 = 10000）
	fmt.Println("Final counter value:", counter)
}
