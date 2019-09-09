package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	preIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[preIndex] {
			// 移动元素值
			preIndex = preIndex + 1
			nums[preIndex] = nums[i]
		}
	}
	return preIndex + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))
}
