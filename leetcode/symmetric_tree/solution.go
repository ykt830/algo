package symmetric_tree

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

提示：
树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return traverse(root.Left, root.Right)
}

func traverse(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}

	if r1.Val != r2.Val {
		return false
	}

	return traverse(r1.Left, r2.Right) && traverse(r1.Right, r2.Left)
}
