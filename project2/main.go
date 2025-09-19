package main

import (
	"fmt"
)

func incrementByTen(numPtr *int) {
	if numPtr == nil {
		fmt.Println("错误：接收到空指针")
		return
	}
	*numPtr += 10
}

func main() {
	value1 := 5
	fmt.Printf("修改前值1: %d\n", value1)
	incrementByTen(&value1)
	fmt.Printf("修改后值1: %d\n", value1) // 预期输出15
}
