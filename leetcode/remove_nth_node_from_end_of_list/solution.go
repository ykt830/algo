package remove_nth_node_from_end_of_list

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


进阶：你能尝试使用一趟扫描实现吗？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, r1 := head, head
	var r2 *ListNode
	count := 1
	for l != nil {
		l = l.Next
		if count > n {
			r2 = r1
			r1 = r1.Next
		}

		count++
	}

	if r2 == nil {
		return head.Next
	}

	r2.Next = r1.Next
	return head
}
