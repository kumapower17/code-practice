package main

func singleNumber(nums []int) int {
	var n = 0
	for _, v := range nums {
		n ^= v
	}
	return n
}
