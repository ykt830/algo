package minimum_size_subarray_sum

import "math"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1

示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

提示：
1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
*/

func minSubArrayLen(target int, nums []int) int {
	i, j, sum, min := 0, 1, nums[0], math.MaxInt
	if sum > target {
		return 1
	}

	for ; j < len(nums); j++ {
		sum += nums[j]
		if sum < target {
			continue
		}

		if j-i+1 < min {
			min = j - i + 1
		}

		for sum-nums[i] >= target {
			sum -= nums[i]
			i++

			if j-i+1 < min {
				min = j - i + 1
			}
		}
	}

	if min == math.MaxInt {
		return 0
	}

	return min
}
