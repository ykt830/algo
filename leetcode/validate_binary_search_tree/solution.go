package validate_binary_search_tree

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

提示：

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate1(root)
}

func validate1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return validate1(root.Left) && validate1(root.Right) && validate2(root)
}

func validate2(root *TreeNode) bool {
	l, r := root.Left, root.Right
	for l != nil {
		if l.Val >= root.Val {
			return false
		}
		l = l.Right
	}
	for r != nil {
		if r.Val <= root.Val {
			return false
		}
		r = r.Left
	}
	return true
}
