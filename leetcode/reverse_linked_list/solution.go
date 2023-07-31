package reverse_linked_list

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var l, r, ret *ListNode = nil, head, head.Next
	for ret != nil {
		r.Next = l
		l = r
		r = ret
		ret = ret.Next
	}
	r.Next = l

	return r
}
