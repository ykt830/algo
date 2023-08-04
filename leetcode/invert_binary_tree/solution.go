package invert_binary_tree

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

提示：
树中节点数目范围在 [0, 100] 内
-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
