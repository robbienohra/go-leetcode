package main

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	p := root

	var res []int

	var travStack []*TreeNode

	for p != nil {
		for p != nil {

			// push nodes onto the stack in reverse access order

			if p.Right != nil {
				travStack = append(travStack, p.Right)
			}

			travStack = append(travStack, p)

			p = p.Left

		}

		// at this point we would have encountered the leftmost left child

		n := len(travStack)

		p = travStack[n-1]

		travStack = travStack[:n-1]

		for len(travStack) > 0 && p.Right == nil {
			res = append(res, p.Val)

			n := len(travStack)

			p = travStack[n-1]

			travStack = travStack[:n-1]
		}

		res = append(res, p.Val)

		n = len(travStack)

		if n > 0 {

			p = travStack[n-1]

			travStack = travStack[:n-1]

		} else {
			p = nil
		}

	}

	return res

}
