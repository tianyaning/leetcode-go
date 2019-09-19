package main

import "fmt"

func twoSum167(numbers []int, target int) []int {
	result := make([]int, 2)
	var tempMap = make(map[int]int)
	for index, value := range numbers {
		lastIndex, ok := tempMap[value]
		if ok {
			result[0] = lastIndex + 1
			result[1] = index + 1
		}
		tempMap[target-value] = index
	}
	return result
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum167(numbers, target))
}
