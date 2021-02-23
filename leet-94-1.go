package main

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, res *[]int) {

	if root != nil {

		// recurse left
		if root.Left != nil {
			helper(root.Left, res)
		}

		// leftmost child will have been encountered at this point
		// push value onto the result set
		*res = append(*res, root.Val)

		// recurse right
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
