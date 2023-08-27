package sliding_window_maximum

import "container/heap"

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v any) {
	*h = append(*h, v.([2]int))
}

func (h *IntHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	h := IntHeap(make([][2]int, 0))
	for i := 0; i < k; i++ {
		h = append(h, [2]int{nums[i], i})
	}

	heap.Init(&h)
	ret := []int{h[0][0]}

	for end := k; end < len(nums); end++ {
		start := end - k + 1
		heap.Push(&h, [2]int{nums[end], end})

		top := h[0]
		for top[1] < start {
			heap.Pop(&h)
			top = h[0]
		}

		ret = append(ret, top[0])
	}

	return ret
}
