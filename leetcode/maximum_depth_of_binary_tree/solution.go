package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树 root ，返回其最大深度。
二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

提示：
树中节点的数量在 [0, 104] 区间内。
-100 <= Node.val <= 100
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1

	if left > right {
		return left
	}

	return right
}
