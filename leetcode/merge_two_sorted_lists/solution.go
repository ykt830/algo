package merge_two_sorted_lists

import "math"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	l1, l2 := list1, list2
	for l1 != nil || l2 != nil {
		a := math.MaxInt
		if l1 != nil {
			a = l1.Val
		}

		b := math.MaxInt
		if l2 != nil {
			b = l2.Val
		}

		var c int
		if a < b {
			c = a
			l1 = l1.Next
		} else {
			c = b
			l2 = l2.Next
		}

		tmp.Next = &ListNode{Val: c}
		tmp = tmp.Next
	}

	return head.Next
}
