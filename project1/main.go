package main

import (
	"fmt"
)

func findSingleNumber(numbers []int) int {
	//创建map存储每个数字出现的次数
	countMap := make(map[int]int)

	//遍历数组，对每个数字进行计数
	for _, number := range numbers {
		countMap[number]++
	}

	//遍历map，找到出现次数为1的数字
	for number, count := range countMap {
		if count == 1 {
			return number
		}
	}

	return -1
}

func main() {
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}

	fmt.Println(findSingleNumber(nums1)) // 输出: 1
	fmt.Println(findSingleNumber(nums2)) // 输出: 4
	fmt.Println(findSingleNumber(nums3)) // 输出: 1
}
