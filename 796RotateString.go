package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {
	//只要看B是不是A+A的子集就行了
	if len(A) != len(B) {
		return false
	}
	return strings.Contains(A+A, B)
}

func main() {
	A := "abcde"
	B := "abced"
	fmt.Println(rotateString(A, B))

}
