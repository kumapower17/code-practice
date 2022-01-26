package main

// https://leetcode-cn.com/problems/search-insert-position/

import "sort"

func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}
