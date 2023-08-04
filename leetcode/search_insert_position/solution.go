package search_insert_position

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。



示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4


提示:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为 无重复元素 的 升序 排列数组
-104 <= target <= 104
*/

func searchInsert(nums []int, target int) int {
	return search(nums, target, 0, len(nums)-1)
}

func search(nums []int, target, i, j int) int {
	if i == j {
		if nums[i] >= target {
			return i
		} else {
			return i + 1
		}
	}

	m := (i + j) / 2
	if nums[m] >= target {
		return search(nums, target, i, m)
	} else {
		return search(nums, target, m+1, j)
	}
}
