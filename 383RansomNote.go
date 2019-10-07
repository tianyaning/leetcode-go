package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	//列出了magazine的字母表，然后算出了出现个数，然后遍历ransomNote，保证有足够的字母可用
	arr := make([]int, 26)
	for _, value := range magazine {
		arr[value-'a']++
	}

	for _, value := range ransomNote {
		arr[value-'a'] = arr[value-'a'] - 1
		if arr[value-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
