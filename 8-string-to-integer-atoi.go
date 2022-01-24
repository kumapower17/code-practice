package main

import "strings"

// https://leetcode-cn.com/problems/string-to-integer-atoi/

const (
	int32Max = (1 << 31) - 1
	digits   = "0123456789"
)

func myAtoi(s string) int {
	cleanStr := strings.TrimLeft(s, " ")
	if cleanStr == "" {
		return 0
	}

	var negative int64
	if cleanStr[0] == '-' {
		negative = 1
		cleanStr = cleanStr[1:]
	} else if cleanStr[0] == '+' {
		cleanStr = cleanStr[1:]
	}

	var val int64
	for _, r := range cleanStr {
		i := strings.IndexByte(digits, byte(r))
		if i == -1 {
			break
		}
		max := int32Max + negative
		if val > (max-int64(i))/10 {
			val = max
			break
		}
		val = 10*val + int64(i)
	}
	if negative == 1 {
		val = -val
	}
	return int(val)
}
