package main

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	var res []int

	var stack []*TreeNode

	curr := root

	for curr != nil || len(stack) > 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// leftmost child would have been encountered at this point
		// pop this child off the stack
		// push value onto the result slice

		n := len(stack)

		curr = stack[n-1]

		stack = stack[:n-1]

		res = append(res, curr.Val)

		// if this is nil (i.e. the leftmost child has no right child)
		// then at the next iteration the initial loop above will be skipped
		// and we simply proceed to access the parent of this current node
		// which is the next leftmost child
		// if this is not nil then at the next iteration we proceed to
		// find the leftmost child of the right subtree

		curr = curr.Right

	}

	return res

}
