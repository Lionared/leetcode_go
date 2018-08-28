package leetcode_go

import (
	"testing"
)

var root *TreeNode

func Test_maxDepth(t *testing.T) {
	depth := maxDepth(root)
	t.Log(depth)
}
