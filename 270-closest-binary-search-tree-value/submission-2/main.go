package main

import "math"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {

	min := math.MinInt64

	var stack []*TreeNode

	curr := root

	for curr != nil || len(stack) > 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		n := len(stack)

		curr = stack[n-1]

		stack = stack[:n-1]

		if math.Abs(float64(curr.Val)-target) < math.Abs(float64(min)-target) {
			min = curr.Val
		}

		curr = curr.Right

	}

	return min

}
