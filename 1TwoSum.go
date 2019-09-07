package main

import "fmt"

func twoSum(nums []int, target int) []int {

	//遍历nums，填充tempMap
	tempMap := make(map[int]int)
	for index, value := range nums {
		tempMap[value] = index
	}

	//定义结果数组
	result := make([]int, 2)

	//再次遍历nums，查找是否有符合条件的
	for index, value := range nums {
		temp := target - value
		tempValue, ok := tempMap[temp]
		if ok && tempValue != index {
			result[0] = index
			result[1] = tempValue
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
