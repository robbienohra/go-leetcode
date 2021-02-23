package main

func productExceptSelf1(nums []int) []int {

	n := len(nums)
	output := make([]int, n)
	L := make([]int, n)
	R := make([]int, n)
	L[0] = 1
	R[n-1] = 1

	for i := 1; i < n; i++ {

		L[i] = L[i-1] * nums[i-1]

	}

	for i := n - 1; i >= 0; i++ {

		R[i] = R[i+1] * nums[i+1]

	}

	for i := 0; i < n; i++ {

		output[i] = L[i] * R[i]

	}

	return output

}
