package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	//将数字转化成字符串，然后首尾对应判断每个字符即可
	stringValue := strconv.Itoa(x)
	for i, j := 0, len(stringValue)-1; i <= j; i, j = i+1, j-1 {
		if stringValue[i] != stringValue[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
