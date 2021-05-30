package main

import "math"

func first(nums []int, target int) int {

	n := len(nums)

	if n == 0 {
		return -1
	}

	L, R := 0, n-1

	for L != R {
		m := int(math.Floor(float64(L) + float64((R-L))/2))

		if nums[m] >= target {
			R = m
		} else {
			L = m + 1
		}

	}

	if nums[L] == target {
		return L
	}

	return -1

}

func last(nums []int, target int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	L, R := 0, n-1

	for L != R {
		m := int(math.Ceil(float64(L) + float64((R-L))/2))

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

func searchRange(nums []int, target int) []int {

	a := first(nums, target)
	b := last(nums, target)

	return []int{a, b}

}
