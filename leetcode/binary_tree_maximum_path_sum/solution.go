package binary_tree_maximum_path_sum

import "math"

/*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

提示：
树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func maxPathSum(root *TreeNode) int {
	ret = math.MinInt
	traverse(root)

	return ret
}

func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := traverse(root.Left), traverse(root.Right)
	v := root.Val
	if left > 0 {
		v += left
	}
	if right > 0 {
		v += right
	}
	ret = max(ret, v)

	return max(root.Val, root.Val+max(left, right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
