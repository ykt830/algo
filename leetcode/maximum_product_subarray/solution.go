package maximum_product_subarray

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32- 位 整数。
子数组是数组的连续子序列。

示例 1:
输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:
输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

提示:
1 <= nums.length <= 2 * 10^4
-10 <= nums[i] <= 10
nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
*/

func maxProduct(nums []int) int {
	ret, minP, maxP := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num < 0 {
			maxP, minP = max(num, minP*num), min(num, maxP*num)
		} else {
			maxP, minP = max(num, maxP*num), min(num, minP*num)
		}

		ret = max(ret, maxP)
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
