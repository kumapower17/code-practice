package main

// https://leetcode-cn.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	first := strs[0]
	remaining := strs[1:]
	for i := 0; i < len(first); i++ {
		for _, s := range remaining {
			if i >= len(s) || s[i] != first[i] {
				return first[:i]
			}
		}
	}
	return first
}
