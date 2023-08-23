package top_k_frequent_elements

import "container/heap"

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]


提示：
1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

进阶：你所设计算法的时间复杂度必须 优于 O(n log n) ，其中 n 是数组大小。
*/

type pair struct {
	k, v int
}

type pairHeap []*pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].v > h[j].v
}

func (h pairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairHeap) Push(e any) {
	*h = append(*h, e.(*pair))
}

func (h *pairHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return v
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	h := new(pairHeap)
	for k_, v_ := range m {
		heap.Push(h, &pair{k: k_, v: v_})
	}

	ret := make([]int, 0)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).(*pair).k)
	}

	return ret
}
