package binary_tree_level_order_traversal

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

提示：
树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		arr: make([]T, 0),
	}
}

func (q *Queue[T]) Push(v T) {
	q.arr = append(q.arr, v)
}

func (q *Queue[T]) Poll() T {
	if len(q.arr) == 0 {
		var zero T
		return zero
	}

	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}

func (q *Queue[T]) Len() int {
	return len(q.arr)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	rets, queue1, queue2 := make([][]int, 0), NewQueue[*TreeNode](), NewQueue[*TreeNode]()
	queue1.Push(root)

	for queue1.Len() > 0 {
		ret := make([]int, 0)
		for queue1.Len() > 0 {
			node := queue1.Poll()
			if node.Left != nil {
				queue2.Push(node.Left)
			}
			if node.Right != nil {
				queue2.Push(node.Right)
			}
			ret = append(ret, node.Val)
		}
		rets = append(rets, ret)
		queue1, queue2 = queue2, queue1
	}

	return rets
}
