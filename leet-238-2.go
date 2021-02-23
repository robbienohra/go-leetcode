package main

func productExceptSelf2(nums []int) []int {

	n := len(nums)
	output := make([]int, n)

	output[0] = 1

	for i := 1; i < n; i++ {

		output[i] = output[i-1] * nums[i-1]

	}

	R := 1

	for i := n - 1; i >= 0; i-- {

		output[i] = output[i] * R

		R *= nums[i]

	}

	return output

}