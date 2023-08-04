package binary_tree_inorder_traversal

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func inorderTraversal(root *TreeNode) []int {
	ret = make([]int, 0)
	if root == nil {
		return ret
	}

	traverse(root)
	return ret
}

func traverse(root *TreeNode) {
	if root.Left != nil {
		traverse(root.Left)
	}

	ret = append(ret, root.Val)

	if root.Right != nil {
		traverse(root.Right)
	}
}
