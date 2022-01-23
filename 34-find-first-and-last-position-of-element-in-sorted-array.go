// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

import "sort"

func searchRange(nums []int, target int) []int {
	l := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}

	rightHalf := nums[l:]
	r := l + sort.Search(len(rightHalf), func(i int) bool {
		return rightHalf[i] > target
	})
	return []int{l, r - 1}
}
