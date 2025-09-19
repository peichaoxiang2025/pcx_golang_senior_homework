package main

import (
	"fmt"
)

func doubleSliceElements(slicePtr *[]int) {
	// 空指针保护
	if slicePtr == nil {
		fmt.Println("错误：接收到空指针")
		return
	}

	// 解引用指针获取切片
	slice := *slicePtr

	// 遍历修改每个元素
	for i := range slice {
		slice[i] *= 2
	}
}

func main() {
	// 测试用例1：正常整数切片
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前: %v\n", numbers)
	doubleSliceElements(&numbers)
	fmt.Printf("修改后: %v\n", numbers) // 预期输出 [2 4 6 8 10]

	// 测试用例2：空切片处理
	emptySlice := []int{}
	fmt.Printf("空切片修改前: %v\n", emptySlice)
	doubleSliceElements(&emptySlice)
	fmt.Printf("空切片修改后: %v\n", emptySlice) // 预期输出 []

	// 测试用例3：负数处理
	negatives := []int{-3, -2, 0, 4}
	fmt.Printf("含负数切片修改前: %v\n", negatives)
	doubleSliceElements(&negatives)
	fmt.Printf("含负数切片修改后: %v\n", negatives) // 预期输出 [-6 -4 0 8]

	// 测试用例4：nil指针处理
	var nilPtr *[]int
	fmt.Println("nil指针处理:")
	doubleSliceElements(nilPtr) // 预期输出错误提示
}
