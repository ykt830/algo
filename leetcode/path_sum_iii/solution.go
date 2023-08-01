package path_sum_iii

/*
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

提示:
二叉树的节点个数的范围是 [0,1000]
-109 <= Node.val <= 109
-1000 <= targetSum <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	hash        map[int]int
	cnt, target int
)

func pathSum(root *TreeNode, targetSum int) int {
	hash, cnt, target = map[int]int{0: 1}, 0, targetSum
	traverse(root, 0)
	return cnt
}

func traverse(root *TreeNode, sum int) {
	if root == nil {
		return
	}

	sum += root.Val
	cnt += hash[sum-target]
	hash[sum]++

	traverse(root.Left, sum)
	traverse(root.Right, sum)

	hash[sum]--
}
