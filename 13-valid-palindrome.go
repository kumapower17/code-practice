package main

import (
	"strings"
	"unicode"
)

// https://leetcode-cn.com/problems/valid-palindrome/

func isAlphanum(b byte) bool {
	const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return strings.IndexByte(alphanum, b) != -1
}

func caseCmp(a, b rune) bool {
	return unicode.ToLower(a) == unicode.ToLower(b)
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if !isAlphanum(s[i]) {
			i++
			continue
		}
		if !isAlphanum(s[j]) {
			j--
			continue
		}
		if !caseCmp(rune(s[i]), rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
