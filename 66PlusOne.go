package main

import "fmt"

func plusOne(digits []int) []int {
	/*
		1.正常数组（末尾位小于9）： 末尾 + 1 返回
		2.末尾位为9，但是不全为9 ：从后向前看，为9的位置置零（不是林志玲），最后一个不为9的位置 + 1 返回
		3.全部为9： 判断位数，返回 位数 + 1 大小的一个数组，首位为1。
	*/
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	arr := make([]int, len(digits)+1)
	arr[0] = 1
	return arr
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(plusOne(nums))
}
