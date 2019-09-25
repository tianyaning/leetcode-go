package main

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	up := int(math.Sqrt(float64(num)))
	for i := 2; i <= up; i++ {
		if (num % i) == 0 {
			sum += i
			sum += num / i
		}
	}

	return sum == num
}
