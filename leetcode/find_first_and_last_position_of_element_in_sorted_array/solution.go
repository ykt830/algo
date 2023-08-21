package find_first_and_last_position_of_element_in_sorted_array

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/

func searchRange(nums []int, target int) []int {
	l, r := findMin(nums, target), findMax(nums, target)
	if l != -1 {
		if nums[l] != target {
			l = -1
		}
	}
	if r != -1 {
		if nums[r] != target {
			r = -1
		}
	}

	return []int{l, r}
}

func findMin(nums []int, target int) int {
	if len(nums) == 0 || nums[0] > target {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m
		} else if nums[m] == target {
			r = m
		} else {
			r = m - 1
		}

		if r-l <= 1 {
			if nums[l] == nums[r] {
				return l
			}

			if target >= nums[r] {
				return r
			} else {
				return l
			}
		}
	}

	return -1
}

func findMax(nums []int, target int) int {
	if len(nums) == 0 || nums[len(nums)-1] < target {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] == target {
			l = m
		} else {
			r = m
		}

		if r-l <= 1 {
			if nums[l] == nums[r] {
				return r
			}

			if target <= nums[l] {
				return l
			} else {
				return r
			}
		}
	}

	return -1
}
