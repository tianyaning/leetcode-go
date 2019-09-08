package main

import "fmt"

func searchInsert(nums []int, target int) int {

	for index, value := range nums {
		if value >= target {
			return index
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	fmt.Println(searchInsert(nums, target))
}
