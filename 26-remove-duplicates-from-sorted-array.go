package main

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/submissions/

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	prev := nums[0]
	n := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[n] = nums[i]
			n++
			prev = nums[i]
		}
	}
	return n
}
