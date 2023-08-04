package diameter_of_binary_tree

/*
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。

提示：
树中节点数目在范围 [1, 104] 内
-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func diameterOfBinaryTree(root *TreeNode) int {
	ret = 0

	traverse(root)
	return ret
}

func traverse(root *TreeNode) int {
	if root == nil {
		return -1
	}

	left := traverse(root.Left)
	right := traverse(root.Right)
	ret = max(ret, left+right+2)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
