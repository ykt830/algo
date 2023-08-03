package kth_largest_element_in_an_array

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：
1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func findKthLargest(nums []int, k int) int {
	begin, end := 0, len(nums)-1
	for {
		mid := partition(nums, begin, end)
		cnt := end - mid + 1
		if cnt == k {
			return nums[mid]
		} else if cnt > k {
			begin = mid + 1
		} else {
			end = mid - 1
			k -= cnt
		}
	}
}

func partition(nums []int, begin, end int) int {
	l, r := begin-1, begin
	for r < end {
		if nums[r] >= nums[end] {
			r++
		} else {
			nums[r], nums[l+1] = nums[l+1], nums[r]

			l++
			r++
		}
	}
	nums[l+1], nums[end] = nums[end], nums[l+1]
	return l + 1
}
