package kth_smallest_element_in_a_bst

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

提示：
树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cnt, ret int

func kthSmallest(root *TreeNode, k int) int {
	cnt, ret = 0, 0
	traverse1(root, k)
	return ret
}

func traverse1(root *TreeNode, k int) {
	if root.Left != nil {
		traverse1(root.Left, k)
	}

	cnt++
	if cnt == k {
		ret = root.Val
	}

	if root.Right != nil {
		traverse1(root.Right, k)
	}
}
