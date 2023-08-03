package construct_binary_tree_from_preorder_and_inorder_traversal

/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
提示:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder 和 inorder 均 无重复 元素
inorder 均出现在 preorder
preorder 保证 为二叉树的前序遍历序列
inorder 保证 为二叉树的中序遍历序列
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	leftInArr, rightInArr := make([]int, 0), make([]int, 0)
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			leftInArr, rightInArr = inorder[:i], inorder[i+1:]
			break
		}
	}

	leftPreArr, rightPreArr := preorder[1:1+len(leftInArr)], preorder[1+len(leftInArr):]
	root.Left = buildTree(leftPreArr, leftInArr)
	root.Right = buildTree(rightPreArr, rightInArr)
	return root
}
