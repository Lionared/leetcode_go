package leetcode_go

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ret []int

	return ret
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int

	return ret
}

func postorderTraversal(root *TreeNode) []int {
	var ret []int

	return ret
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int

	return ret
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	height := maxInt(left, right)
	return height
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
