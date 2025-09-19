package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型：无参数、返回结果和错误的函数
type Task func() (interface{}, error)

// TaskResult 存储任务执行结果的详细信息
type TaskResult struct {
	TaskID   int           // 任务唯一标识（索引）
	Start    time.Time     // 任务开始时间
	End      time.Time     // 任务结束时间
	Duration time.Duration // 任务执行耗时
	Result   interface{}   // 任务返回结果（成功时有效）
	Error    error         // 任务执行错误（失败时有效）
}

// Scheduler 任务调度器核心结构体
type Scheduler struct {
	tasks       []Task          // 待执行的任务列表
	concurrency int             // 最大并发数（同时运行的任务数）
	results     chan TaskResult // 结果输出通道（缓冲大小等于任务数）
	wg          sync.WaitGroup  // 用于等待所有任务完成
}

// NewScheduler 创建新调度器实例
func NewScheduler(tasks []Task, concurrency int) *Scheduler {
	return &Scheduler{
		tasks:       tasks,
		concurrency: concurrency,
		results:     make(chan TaskResult, len(tasks)), // 缓冲通道避免阻塞
	}
}

// Run 启动任务调度器，执行所有任务并收集结果
func (s *Scheduler) Run() {
	defer close(s.results) // 确保所有任务完成后关闭结果通道

	semaphore := make(chan struct{}, s.concurrency) // 控制并发数的信号量

	for i, task := range s.tasks {
		s.wg.Add(1)
		go func(taskID int, t Task) {
			defer s.wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			semaphore <- struct{}{} // 获取信号量（阻塞直到有可用槽位）

			start := time.Now()
			result, err := t() // 执行任务
			end := time.Now()

			// 构造并发送结果到通道
			s.results <- TaskResult{
				TaskID:   taskID,
				Start:    start,
				End:      end,
				Duration: end.Sub(start),
				Result:   result,
				Error:    err,
			}
		}(i, task)
	}

	s.wg.Wait() // 等待所有任务完成
}

// 示例用法
func main() {
	// 1. 定义测试任务
	tasks := []Task{
		func() (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return "任务1完成", nil
		},
		func() (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return "任务2完成", nil
		},
		func() (interface{}, error) {
			time.Sleep(150 * time.Millisecond)
			return nil, fmt.Errorf("任务3执行失败")
		},
		func() (interface{}, error) {
			time.Sleep(300 * time.Millisecond)
			return 42, nil
		},
	}

	// 2. 创建调度器（最大并发数 2）
	scheduler := NewScheduler(tasks, 2)

	// 3. 启动调度器（在后台运行）
	go func() {
		scheduler.Run()
	}()

	// 4. 等待调度器完成并收集结果
	var results []TaskResult
	for result := range scheduler.results {
		results = append(results, result)
	}

	// 5. 打印所有任务执行结果
	fmt.Println("任务执行结果：")
	for _, result := range results {
		fmt.Printf(
			"任务ID: %-2d | 开始: %-20s | 结束: %-20s | 耗时: %-6v | 结果: %-20v | 错误: %v\n",
			result.TaskID,
			result.Start.Format("15:04:05.000"),
			result.End.Format("15:04:05.000"),
			result.Duration,
			result.Result,
			result.Error,
		)
	}

	// 6. 统计汇总（关键修复：基于实际结果统计）
	totalTasks := len(results)
	successCount := 0
	totalDuration := time.Duration(0)

	for _, result := range results {
		if result.Error == nil {
			successCount++
		}
		totalDuration += result.Duration
	}

	avgDuration := time.Duration(0)
	if totalTasks > 0 {
		avgDuration = totalDuration / time.Duration(totalTasks)
	}

	fmt.Printf("\n汇总：总任务数 %d，成功 %d，平均耗时 %v\n", totalTasks, successCount, avgDuration)
}
