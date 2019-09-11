package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	kv := strings.Split(strings.TrimSpace(s), " ")
	return len(kv[len(kv)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("a "))
}
