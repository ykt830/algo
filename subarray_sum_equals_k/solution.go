package subarray_sum_equals_k

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2

提示：

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

func subarraySum(nums []int, k int) int {
	i, sum, m, ret := 0, 0, map[int]int{0: 1}, 0
	for ; i < len(nums); i++ {
		sum += nums[i]
		ret += m[sum-k]
		m[sum]++
	}

	return ret
}
