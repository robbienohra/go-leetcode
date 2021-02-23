package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, res *[]int) {

	if root != nil {

		if root.Left != nil {
			helper(root.Left, res)
		}

		*res = append(*res, root.Val)

		if root.Right != nil {
			helper(root.Right, res)
		}
	}

}

func inorderTraversal(root *TreeNode) []int {

	var res []int

	helper(root, &res)

	return res

}

func closestValue(root *TreeNode, target float64) int {

	res := inorderTraversal(root)

	min := res[0]

	for _, val := range res {
		if math.Abs(float64(val)-target) < math.Abs(float64(min)-target) {
			min = val
		}
	}

	return min

}
