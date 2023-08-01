package lowest_common_ancestor_of_a_binary_tree

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

提示：

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	parents map[*TreeNode]*TreeNode
	visits  map[*TreeNode]bool
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parents, visits = make(map[*TreeNode]*TreeNode), make(map[*TreeNode]bool)

	parents[root] = nil
	traverse(root)

	n1, n2 := p, q

	visits[n1] = true
	for parents[n1] != nil {
		n1 = parents[n1]
		visits[n1] = true
	}

	for {
		if visits[n2] {
			return n2
		}
		n2 = parents[n2]
	}
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		parents[root.Left] = root
	}
	if root.Right != nil {
		parents[root.Right] = root
	}

	traverse(root.Left)
	traverse(root.Right)
}
