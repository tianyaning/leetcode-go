package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func shortestToChar(S string, C byte) []int {
	n := len(S)
	pos := -n
	ans := make([]int, len(S))
	for i := 0; i < n; i++ {
		if S[i] == C {
			pos = i
		}
		ans[i] = i - pos
	}
	for i := n - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		ans[i] = min(ans[i], abs(i-pos))
	}
	return ans
}
