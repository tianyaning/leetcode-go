package main

import "strings"

func isPalindrome(s string) bool {
	if len(s) <= 0 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1

	for i <= j {
		if isValidCharacter(s[i]) && isValidCharacter(s[j]) && (s[i] == s[j]) {
			i++
			j--
			continue
		} else if isValidCharacter(s[i]) && isValidCharacter(s[j]) && (s[i] != s[j]) {
			return false
		}
		if !isValidCharacter(s[i]) {
			i++
		}
		if !isValidCharacter(s[j]) {
			j--
		}
	}
	return true
}

func isValidCharacter(b byte) bool {
	if ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}
