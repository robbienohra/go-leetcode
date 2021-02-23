package main

func search(nums []int, target int) int {

	n := len(nums)

	if n == 0 {
		return -1
	}

	L := 0
	R := n - 1

	for L <= R {
		m := mid1(L, R)

		if nums[m] < target {
			L = m + 1
		} else if nums[m] > target {
			R = m - 1
		} else {
			return m
		}

	}

	return -1

}

func mid(L, R int) int {
	return L + (R-L)/2
}
