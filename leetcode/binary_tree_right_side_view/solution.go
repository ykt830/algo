package binary_tree_right_side_view

/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

提示:
二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	maxLevel int
	ret      []int
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	maxLevel, ret = -1, []int{}
	traverse(root, 0)
	return ret
}

func traverse(root *TreeNode, level int) {
	if level > maxLevel {
		ret = append(ret, root.Val)
		maxLevel = level
	}

	if root.Right != nil {
		traverse(root.Right, level+1)
	}
	if root.Left != nil {
		traverse(root.Left, level+1)
	}
}
