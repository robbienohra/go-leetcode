package main

import "math"

func search(nums []int, target int) int {

	n := len(nums)

	if n == 0 {
		return -1
	}

	L := 0
	R := n - 1

	for L != R {
		m := mid(L, R)

		if nums[m] <= target {
			L = m
		} else {
			R = m - 1
		}

	}

	if nums[L] == target {
		return L
	}

	return -1

}

func mid(L, R int) int {
	return L + int(math.Ceil(float64((R-L))/2))
}
