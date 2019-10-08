package main

import "fmt"

func largeGroupPositions(S string) [][]int {
	var rs [][]int
	var i int
	for i < len(S) {
		s := i
		b := S[i]
		for i < len(S) && S[i] == b {
			i++
		}
		if i-s >= 3 {
			rs = append(rs, []int{s, i - 1})
		}
	}
	return rs
}

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abc"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
}
