package merge_k_sorted_lists

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var ret = &ListNode{}

func mergeKLists(lists []*ListNode) *ListNode {
	heads := make([]*ListNode, 0)
	for _, node := range lists {
		heads = append(heads, node)
	}

	tail := ret
	for {
		pos := merge(heads)
		if pos == -1 {
			break
		}
		tail.Next = heads[pos]
		tail = tail.Next
		heads[pos] = heads[pos].Next
	}

	return ret.Next
}

func merge(heads []*ListNode) int {
	var minNode *ListNode
	var pos = -1
	for i, head := range heads {
		if head == nil {
			continue
		}

		if minNode == nil {
			minNode = head
			pos = i
			continue
		}

		if head.Val < minNode.Val {
			minNode = head
			pos = i
		}
	}

	return pos
}
